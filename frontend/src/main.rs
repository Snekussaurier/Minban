#![allow(non_snake_case)]

use api::*;
use components::icons::Plus;
use components::*;
use dioxus::logger::tracing::{error, info};
use dioxus::prelude::*;
use mods::*;
use reqwest::Error;
use std::collections::{BTreeSet, HashMap};
use std::result::Result::Ok;
use utils::{IsNewCardState, IsSelectingState, LoginState};

mod api;
mod components;
mod mods;
mod utils;

#[derive(Clone, Routable, Debug, PartialEq)]
enum Route {
    #[route("/")]
    Dashboard {},
}

fn main() {
    launch(App);
}

fn App() -> Element {
    let login_state = use_resource(check_auth);
    const VERSION: &str = env!("CARGO_PKG_VERSION");
    rsx! {
        document::Stylesheet {
            // Urls are relative to your Cargo.toml file
            href: asset!("assets/tailwind.css")
        }
        div {
            class: "flex items-center justify-center h-screen w-screen bg-gray-100 relative",
            match &*login_state.read_unchecked() {
                Some(Ok(LoginState::NotLoggedIn)) => rsx!( Login {
                    login_state
                }
                p {
                    class: "absolute right-1 bottom-1 text-sm text-slate-400",
                    "Minban {VERSION}"
                }
                ),
                Some(Ok(LoginState::LoggedIn)) => rsx!( Dashboard {} ),
                Some(Err(_)) =>  rsx!{ Loading {} },
                None => rsx!{ Loading {} },
            }
        }
    }
}

#[component]
fn Dashboard() -> Element {
    let mut boards_signal = use_context_provider(|| Signal::new(BoardLeanModel::default()));
    let mut board_signal = use_context_provider(|| Signal::new(HashMap::new()));
    let mut states_signal = use_context_provider(|| Signal::new(Vec::new() as Vec<StateModel>));
    let mut tags_signal = use_context_provider(|| Signal::new(Vec::new() as Vec<TagModel>));

    let board = use_resource(move || async move {
        let fetched_data = fetch_data().await;
        match fetched_data {
            Ok(fetched_data) => {
                boards_signal.set(fetched_data.board);
                board_signal.set(fetched_data.cards);
                states_signal.set(fetched_data.columns);
                tags_signal.set(fetched_data.tags);
                Ok(())
            }
            Err(err) => Err(err),
        }
    });

    let _selected_card_context = use_context_provider(|| Signal::new(CardModel::default()));
    let _is_selecting = use_context_provider(|| Signal::new(IsSelectingState(false)));
    let _is_new_card = use_context_provider(|| Signal::new(IsNewCardState(false)));

    let mut settings_open = use_signal(|| false);

    match &*board.read_unchecked() {
        Some(Ok(_)) => {
            rsx! (
                div{
                    class: "flex flex-col h-full w-full py-6 bg-white",
                    Header {
                        on_click_settings: move |_| {
                            settings_open.set(true);
                        }
                    }
                    CardDetails {
                        on_create: move |mut card: CardModel| {
                            spawn(async move {
                                let created_card_id = create_card(&boards_signal().id, &card).await;
                                match created_card_id {
                                    Ok(id_str) => {
                                        card.id = id_str;
                                        board_signal.write().entry(card.state_id).or_insert_with(BTreeSet::new).insert(card);
                                    }
                                    Err(err) => {
                                        error!("Error creating card: {:?}", err);
                                    }
                                }
                            });
                        },
                        on_update: move |card: CardModel| {
                            let mut board = board_signal.write();
                            let column = board.entry(card.state_id).or_insert_with(BTreeSet::new);

                            column.replace(card.clone());
                            spawn(async move {
                                let updated_card = patch_card(&boards_signal().id, &card).await;

                                match updated_card {
                                    Ok(_) => {}
                                    Err(err) => {
                                        error!("Error updating card {:?}", err);
                                    }
                                }
                            });
                        },
                        on_delete: move |card: CardModel| {
                            let mut board = board_signal.write();
                            let cards = board.entry(card.state_id).or_insert_with(BTreeSet::new);
                            cards.remove(&card);

                            spawn(async move {
                                let deleted_card = delete_card(&boards_signal().id, card.id).await;
                                match deleted_card {
                                    Ok(_) => {}
                                    Err(err) => {
                                        error!("Error updating card {:?}", err);
                                    }
                                }
                            });
                        }
                    }
                    ColumnContainer{
                        states: states_signal(),
                        board: board_signal()
                    }
                    if settings_open() {
                        Settings {
                            on_click_close: move || {
                                settings_open.set(false);
                            }
                        }
                    }
                }
            )
        }
        Some(Err(err)) => {
            rsx! { div { "Error loading board {err}" } }
        }
        _ => {
            rsx! { Loading {} }
        }
    }
}

#[component]
fn ColumnContainer(states: Vec<StateModel>, board: HashMap<u32, BTreeSet<CardModel>>) -> Element {
    rsx!(
        div {
            class: "flex flex-row w-full h-full gap-4 mt-4 overflow-x-auto overflow-y-hidden rounded-md px-6",
            for state in states {
                Column {
                state: state.clone(),
                    cards: {
                        let cards_option = board.get(&state.id);
                        match cards_option {
                            Some(cards) => cards.clone(),
                            None => BTreeSet::new()
                        }
                    },
                }
            }
            button {
                class: "rounded-md bg-slate-100 text-slate-400 hover:text-[#413a46] duration-200 flex items-center justify-center h-12 min-w-12",
                onclick: move |_| {
                    info!("Add state");
                },
                Plus {}
            }
        }
    )
}

async fn fetch_data() -> Result<FetchResponse, Error> {
    Ok(get_board().await?)
}
