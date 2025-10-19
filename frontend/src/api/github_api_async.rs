use crate::mods::GitHubRelease;

pub async fn get_latest_release() -> Result<GitHubRelease, reqwest::Error> {
    let client = reqwest::Client::builder().build()?;

    let url = "https://api.github.com/repos/Snekussaurier/minban/releases/latest";

    let response = client
        .get(url)
        .header(
            "User-Agent",
            &format!("MyDioxusApp/{}", env!("CARGO_PKG_VERSION")),
        )
        .send()
        .await?
        .error_for_status()?;

    Ok(response.json().await?)
}
