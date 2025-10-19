use crate::api::{get_latest_release, patch_board, patch_states, patch_tags};
use crate::icons::X;
use crate::mods::{BoardLeanModel, StateModel};
use crate::TagModel;
use dioxus::prelude::*;
use serde_json::value;

#[component]
pub fn Settings(on_click_close: EventHandler) -> Element {
    let current_setting_page = use_signal(|| SettingsView::Board);

    let mut board_signal = use_context::<Signal<BoardLeanModel>>();
    let mut columns = use_context::<Signal<Vec<StateModel>>>();
    let mut tags = use_context::<Signal<Vec<TagModel>>>();

    let draft_board = use_signal(|| board_signal.read().clone());
    let draft_columns = use_signal(|| columns.read().clone());
    let draft_tags = use_signal(|| tags.read().clone());

    let mut save_settings = move || {
        board_signal.set(draft_board());
        columns.set(draft_columns.read().to_vec());
        tags.set(draft_tags.read().to_vec());
        spawn(async move {
            patch_board(&board_signal()).await;
            patch_states(&board_signal().id, &columns()).await;
            patch_tags(&board_signal().id, &tags()).await;

            on_click_close.call(());
        });
    };

    rsx! {
        div {
            class: "absolute top-0 left-0 z-50 w-full h-full flex backdrop-blur-sm items-center justify-center bg-slate-950/20",
            div {
                class: "max-w-5xl w-full max-h-[35rem] h-full bg-white rounded-md flex flex-row overflow-hidden",
                SettingsNavigation {
                    current_setting_page
                }
                div {
                    class: "w-full h-full flex flex-col",
                    div {
                        class: "w-full flex flex-row justify-between border-b border-slate-100 p-4",
                        h1 {
                            class: "text-2xl text-minban-dark",
                            {current_setting_page.read().title()}
                        }
                        button {
                            class: "text-slate-400 hover:text-minban-dark duration-200",
                            onclick: move |_| {
                                on_click_close.call(());
                            },
                            X {}
                        }
                    }
                    div {
                        class: "p-4 overflow-auto w-full h-full",
                        match *current_setting_page.read() {
                            SettingsView::Board => rsx! { BoardSettings { board: draft_board } },
                            SettingsView::Column => rsx! { ColumnSettings { columns: draft_columns } },
                            SettingsView::Tag => rsx! { TagSettings { tags: draft_tags } },
                            SettingsView::SoftwareUpdates => rsx! { SoftwareUpdatesSettings {  } }
                        }
                    }
                    div {
                        class: "w-full flex flex-row gap-4 border-t border-slate-100 p-4 justify-end",
                        button {
                            class: "rounded-md w-32 p-2 bg-minban-dark hover:bg-minban-highlight text-white duration-200",
                            onclick: move |_| {
                                save_settings();
                            },
                            "Save"
                        }
                        button {
                            class: "rounded-md w-32 p-2 bg-slate-100 text-slate-400 hover:text-minban-dark duration-200",
                            onclick: move |_| {
                                on_click_close.call(());
                            },
                            "Close"
                        }

                    }
                }
            }
        }
    }
}

#[component]
fn SettingsNavigation(mut current_setting_page: Signal<SettingsView>) -> Element {
    rsx! {
        div {
            class: "h-full w-80 bg-slate-100 flex flex-col p-4 gap-2" ,
            h1 {
                class: "mb-4",
                "Settings"
            }
            NavigationButton {
                text: "Board".to_string(),
                current_selection: current_setting_page(),
                click_value: SettingsView::Board,
                on_click: move |_| {
                    current_setting_page.set(SettingsView::Board);
                }
            }
            NavigationButton {
                text: "Columns".to_string(),
                current_selection: current_setting_page(),
                click_value: SettingsView::Column,
                on_click: move |_| {
                    current_setting_page.set(SettingsView::Column);
                }
            }
            NavigationButton {
                text: "Tags".to_string(),
                current_selection: current_setting_page(),
                click_value: SettingsView::Tag,
                on_click: move |_| {
                    current_setting_page.set(SettingsView::Tag);
                }
            }
            NavigationButton {
                text: "Software Updates".to_string(),
                current_selection: current_setting_page(),
                click_value: SettingsView::SoftwareUpdates,
                on_click: move |_| {
                    current_setting_page.set(SettingsView::SoftwareUpdates);
                }
            }
        }
    }
}

#[component]
fn NavigationButton(
    text: String,
    current_selection: SettingsView,
    click_value: SettingsView,
    on_click: EventHandler,
) -> Element {
    let button_style = if current_selection == click_value {
        "bg-minban-dark text-white"
    } else {
        "text-slate-400 hover:bg-slate-200 hover:text-minban-dark"
    };

    rsx! {
        button {
            style: button_style,
            class: "rounded-md py-2 px-4 {button_style} text-left duration-200",
            onclick: move |_| {
                on_click.call(());
            },
            {text}
        }
    }
}

#[component]
fn BoardSettings(board: Signal<BoardLeanModel>) -> Element {
    rsx! {
        div {
            class: "flex flex-col space-y-2",
            div {
                class: "flex flex-col",
                p {
                    class: "text-slate-400 mb-2",
                    "Board Name"
                }
                input {
                    class: "w-full text-slate-900 text-lg bg-transparent border-none placeholder-slate-400",
                    value: board().name,
                    maxlength: 20,
                    placeholder: "Enter board name",
                    oninput: move |evt| board.write().name = evt.value()
                }
                p {
                    class: "text-xs text-slate-400 min-w-0 mt-1",
                    "Max 20 characters"
                }
            }

            div {
                class: "flex flex-col mt-2 pt-2 border-t border-slate-100",
                p {
                    class: "text-slate-400 mb-2",
                    "Board Description"
                }
                input {
                    class: "w-full text-slate-900 text-lg bg-transparent border-none placeholder-slate-400",
                    value: board().description,
                    maxlength: 120,
                    placeholder: "Describe what this board is about",
                    oninput: move |evt| board.write().description = evt.value()
                }
                p {
                    class: "text-xs text-slate-400 min-w-0 mt-1",
                    "Max 120 characters"
                }
            }

            div {
                class: "flex flex-col mt-2 pt-2 border-t border-slate-100",
                p {
                    class: "text-slate-400 mb-2",
                    "Token"
                }
                input {
                    class: "w-full text-slate-900 text-lg bg-transparent border-none placeholder-slate-400 font-mono",
                    value: board().token,
                    maxlength: 4,
                    placeholder: "ABC",
                    oninput: move |evt| board.write().token = evt.value()
                }
                p {
                    class: "text-xs text-slate-400 min-w-0 mt-1",
                    "Card prefix (max 4)"
                }
            }
        }
    }
}

#[component]
fn ColumnSettings(columns: Signal<Vec<StateModel>>) -> Element {
    let mut edit_column_name = move |id: usize, new_name: String| {
        columns.with_mut(|column_vec| {
            if id < column_vec.len() {
                column_vec[id].name = new_name;
            }
        })
    };

    let mut edit_column_color = move |id: usize, new_color: String| {
        columns.with_mut(|column_vec| {
            if id < column_vec.len() && is_valid_color(&new_color) {
                column_vec[id].color = new_color;
            }
        })
    };

    rsx! {
        table {
            class: "w-full",
            thead {
                class: "text-slate-400",
                tr {
                    td {
                        "Name"
                    }
                    td {
                        "Color"
                    }
                    td {}
                }
            }
            tbody {
                for (id, column) in columns().iter().enumerate() {
                    tr {
                        class: "border-b border-slate-100",
                        td {
                            class: "py-2 pr-2 w-7/12",
                            input {
                                class: "w-full",
                                value: column.name.clone(),
                                maxlength: 20,
                                oninput: move |evt| edit_column_name(id, evt.value())
                            }
                        }
                        td {
                            class: "py-2 w-4/12",
                            ColorValidatorInput {
                                initial_value: column.color.clone(),
                                on_input: move |new_color: String| {
                                    if is_valid_color(&new_color) { edit_column_color(id, new_color) };
                                }
                            }

                        }
                        td {
                            class: "py-2 w-1/12",
                            SettingsColorPreview {
                                color: column.color.clone()
                            }
                        }
                    }
                }
            }
        }
    }
}

#[component]
fn TagSettings(tags: Signal<Vec<TagModel>>) -> Element {
    let mut edit_tag_name = move |id: usize, new_name: String| {
        tags.with_mut(|tag_vec| {
            if id < tag_vec.len() {
                tag_vec[id].name = new_name;
            }
        })
    };

    let mut edit_tag_color = move |id: usize, new_color: String| {
        if !is_valid_color(&new_color) {
            return;
        }
        tags.with_mut(|tag_vec| {
            if id < tag_vec.len() {
                tag_vec[id].color = new_color;
            }
        })
    };

    rsx! {
        table {
            class: "w-full",
            thead {
                class: "text-slate-400",
                tr {
                    td {
                        "Name"
                    }
                    td {
                        "Color"
                    }
                    td {}
                }
            }
            tbody {
                for (id, tag) in tags().iter().enumerate() {
                    tr {
                        class: "border-b border-slate-100",
                        td {
                            class: "py-2 pr-2 w-7/12",
                            input {
                                class: "w-full",
                                value: tag.name.clone(),
                                maxlength: 20,
                                oninput: move |evt| edit_tag_name(id, evt.value())
                            }
                        }
                        td {
                            class: "py-2 w-4/12",
                            ColorValidatorInput {
                                initial_value: tag.color.clone(),
                                on_input: move |new_color: String| {
                                    if is_valid_color(&new_color) { edit_tag_color(id, new_color) };
                                }
                            }
                        }
                        td {
                            class: "py-2 w-1/12",
                            SettingsColorPreview {
                                color: tag.color.clone()
                            }
                        }
                    }
                }
            }
        }
    }
}

#[component]
fn SoftwareUpdatesSettings() -> Element {
    let latest = use_resource(get_latest_release);

    rsx! {
        match &*latest.read_unchecked() {
            Some(Ok(latest_tag)) => {
                let current_version = env!("CARGO_PKG_VERSION");
                let status = compare_versions(current_version, &latest_tag.tag_name);

                rsx! {
                    p {
                        "Current version: v{current_version}"
                    }
                    p {
                        "Latest version: {latest_tag.tag_name}"
                    }
                    p {
                        match status {
                            Ok(std::cmp::Ordering::Less) => "Update available!",
                            Ok(std::cmp::Ordering::Equal) => "You're up to date",
                            Ok(std::cmp::Ordering::Greater) => "You're running a pre-release version",
                            Err(_) => "Version comparison failed"
                        }
                    }
                    textarea {
                        class: "w-full h-full font-mono",
                        value: "{latest_tag.body}",
                        resize: "None"

                    }
                }
            },
            None => rsx!{ "Loading..." },
            Some(Err(err)) => rsx! { "{err}" }
        }
    }
}

#[derive(PartialEq, Clone)]
enum SettingsView {
    Board,
    Column,
    Tag,
    SoftwareUpdates,
}

fn compare_versions(current: &str, latest: &str) -> Result<std::cmp::Ordering, semver::Error> {
    use semver::Version;

    let current_clean = current.trim_start_matches('v');
    let latest_clean = latest.trim_start_matches('v');

    let current_version = Version::parse(current_clean)?;
    let latest_version = Version::parse(latest_clean)?;

    Ok(current_version.cmp(&latest_version))
}

trait SettingsTrait {
    fn title(&self) -> &str;
}

impl SettingsTrait for SettingsView {
    fn title(&self) -> &str {
        match self {
            SettingsView::Board => "Board",
            SettingsView::Column => "Columns",
            SettingsView::Tag => "Tags",
            SettingsView::SoftwareUpdates => "Software Updates",
        }
    }
}

fn is_valid_color(color: &str) -> bool {
    if color.len() != 6 {
        return false;
    }

    color.chars().all(|c| c.is_ascii_hexdigit())
}

#[component]
fn ColorValidatorInput(initial_value: String, on_input: EventHandler<String>) -> Element {
    let mut input_value = use_signal(|| initial_value);

    rsx! {
        input {
            value: input_value,
            maxlength: 6,
            oninput: move |evt| {
                input_value.set(evt.value().clone());
                on_input.call(evt.value());
            }
        }
    }
}

#[component]
fn SettingsColorPreview(color: String) -> Element {
    rsx! {
        div {
            style: "background-color: #{color}",
            class: "rounded-full w-4 h-4"
        }
    }
}
