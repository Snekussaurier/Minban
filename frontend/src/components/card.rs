use dioxus::prelude::*;

use crate::components::Tag;
use crate::mods::{CardModel, TagModel};
use crate::utils::{IsNewCardState, IsSelectingState};

#[component]
pub fn Card(card: CardModel) -> Element {
    let mut selected_card = use_context::<Signal<CardModel>>();
    let mut is_selecting = use_context::<Signal<IsSelectingState>>();
    let mut is_new_card = use_context::<Signal<IsNewCardState>>();
    let tags = use_context::<Signal<Vec<TagModel>>>();

    rsx! {
        div {
            draggable: "true",
            onclick: {
                let card = card.clone();
                move |_| {
                    // Open card details window
                    selected_card.set(card.clone());
                    is_new_card.set(IsNewCardState(false));
                    is_selecting.set(IsSelectingState(true));
                }
            },
            ondrag: {
                let card = card.clone();
                move |_| {
                    selected_card.set(card.clone());
                }
            },
            class: "min-h-24 w-full bg-white hover:bg-slate-100 transition-color duration-200 rounded-md p-4 shrink-0 flex flex-col gap-2 cursor-pointer relative",
            h1 { "{card.title}" }
            div {
                class: "flex flex-row flex-wrap gap-2 grow items-end",
                for tag_id in card.tags.iter() {
                    if let Some(tag) = get_tag_by_id(*tag_id, &tags.read()) {
                        // Or if using HashMap: if let Some(tag) = tag_map.get(tag_id) {
                        Tag { name: tag.name.clone(), color: tag.color.clone() }
                    }
                }
            }
        }
    }
}

fn get_tag_by_id(id: u32, all_tags: &[TagModel]) -> Option<&TagModel> {
    all_tags.iter().find(|tag| tag.id == id)
}
