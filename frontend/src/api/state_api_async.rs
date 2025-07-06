use crate::api::routes::{API_VERSION, BASE_API_URL};
use crate::mods::IdResponse;
use crate::mods::StateModel;
use dioxus::logger::tracing::debug;

pub async fn create_state(board_id: &String, state: StateModel) -> Result<u32, reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!("{}{}{}{}", BASE_API_URL, API_VERSION, board_id, "state");

    let response = client
        .post(url)
        .json(&state)
        .header("Content-Type", "application/json")
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    let id_response: IdResponse = response.json().await?;

    Ok(id_response.id)
}

pub async fn patch_state(board_id: &String, state: StateModel) -> Result<(), reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!(
        "{}{}{}{}{}",
        BASE_API_URL, API_VERSION, board_id, "/state/", &state.id
    );

    let _response = client
        .patch(url)
        .json(&state)
        .header("Content-Type", "application/json")
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    Ok(())
}

pub async fn patch_states(
    board_id: &String,
    states: &Vec<StateModel>,
) -> Result<(), reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!("{}{}{}{}", BASE_API_URL, API_VERSION, board_id, "/states");

    let _response = client
        .patch(url)
        .json(&states)
        .header("Content-Type", "application/json")
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    Ok(())
}

pub async fn delete_state(board_id: &String, state: StateModel) -> Result<(), reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!(
        "{}{}{}{}{}",
        BASE_API_URL, API_VERSION, board_id, "/state/", &state.id
    );

    let _response = client
        .delete(url)
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    Ok(())
}
