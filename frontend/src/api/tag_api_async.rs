use crate::api::routes::{API_VERSION, BASE_API_URL};
use crate::mods::IdResponse;
use crate::mods::TagModel;

pub async fn create_tag(board_id: &String, tag: TagModel) -> Result<u32, reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!("{}{}{}{}", BASE_API_URL, API_VERSION, board_id, "/tag");

    let response = client
        .post(url)
        .json(&tag)
        .header("Content-Type", "application/json")
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    let id_response: IdResponse = response.json().await?;

    Ok(id_response.id)
}

pub async fn patch_tag(board_id: &String, tag: TagModel) -> Result<(), reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!(
        "{}{}{}{}{}",
        BASE_API_URL, API_VERSION, board_id, "/tag/", &tag.id
    );

    let _response = client
        .patch(url)
        .json(&tag)
        .header("Content-Type", "application/json")
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    Ok(())
}

pub async fn patch_tags(board_id: &String, tags: &Vec<TagModel>) -> Result<(), reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!("{}{}{}{}", BASE_API_URL, API_VERSION, board_id, "/tags");

    let response = client
        .patch(url)
        .json(tags)
        .header("Content-Type", "application/json")
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    Ok(())
}

pub async fn delete_tag(board_id: &String, tag: TagModel) -> Result<(), reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = format!(
        "{}{}{}{}{}",
        BASE_API_URL, API_VERSION, board_id, "/tag/", &tag.id
    );

    let _response = client
        .delete(url)
        .header("Content-Type", "application/json")
        .fetch_credentials_include()
        .send()
        .await?
        .error_for_status()?;

    Ok(())
}
