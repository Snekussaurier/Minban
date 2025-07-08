use std::collections::BTreeSet;
use std::collections::HashMap;

use dioxus::logger::tracing::debug;

use crate::api::routes::{API_VERSION, BASE_API_URL};

use crate::mods::FetchResponse;
use crate::mods::{BoardLeanModel, BoardModel, CardModel};

pub async fn get_board() -> Result<FetchResponse, reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!("{}{}{}", BASE_API_URL, API_VERSION, "board");

    let response = client
        .get(url)
        .header("Content-Type", "application/json")
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    let board: BoardModel = response.json().await?;

    let mut board_cards: HashMap<u32, BTreeSet<CardModel>> = HashMap::new();

    for card in board.cards {
        board_cards
            .entry(card.state_id)
            .or_insert_with(BTreeSet::new)
            .insert(card);
    }

    let fetch_response = FetchResponse {
        board: BoardLeanModel {
            id: board.id,
            name: board.name,
            description: board.description,
            token: board.token,
            selected: board.selected,
        },
        cards: board_cards,
        columns: board.states,
        tags: board.tags,
    };

    Ok(fetch_response)
}

pub async fn patch_board(board: &BoardLeanModel) -> Result<(), reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!("{}{}{}{}", BASE_API_URL, API_VERSION, "board/", board.id);

    let response = client
        .patch(url)
        .json(board)
        .header("Content-Type", "application/json")
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    Ok(())
}
