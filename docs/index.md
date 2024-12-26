---
hide:
    - toc
---

<div style="text-align: center;">
  <img src="img/minban-header.png" alt="Task Board" width="400"/>
</div>

<br>

<div style="text-align: center;">
    <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/Snekussaurier/minban-backend/build-and-publish.yaml?style=for-the-badge">
    <img alt="GitHub License" src="https://img.shields.io/github/license/Snekussaurier/minban-backend?style=for-the-badge">
    <img alt="GitHub Issues or Pull Requests" src="https://img.shields.io/github/issues/Snekussaurier/minban-backend?style=for-the-badge">
    <img alt="GitHub Release" src="https://img.shields.io/github/v/release/Snekussaurier/minban-backend?sort=semver&display_name=tag&style=for-the-badge">
</div>

<br>

## Welcome to the **MinBan API Documentation**!  

MinBan is a **minimal Kanban application** designed to help you manage tasks for your projects efficiently. 
The backend is powered by **Go** with the **Gin framework**, offering a fast and lightweight API to handle all your Kanban-related needs. 

## Table of Contents

- **[Installation](installation/index.md)**: Set up MinBan locally.
  - [Docker Setup](installation/docker.md)
  - [Source Setup](installation/source.md)
- **API Routes**:
  - **Authentication**: Learn how to authenticate with the API and perform login, logout, and session checks.
    - [Login](routes/authentication/login.md)
    - [Logout](routes/authentication/logout.md)
    - [Check Auth](routes/authentication/check_auth.md)
  - **Card**: Manage cards within a project.
    - [Get Cards](routes/card/get_cards.md)
    - [Create Card](routes/card/create_card.md)
    - [Update Card](routes/card/update_card.md)
    - [Delete Card](routes/card/delete_card.md)
  - **State**: Manage task states (or columns).
    - [Get States](routes/state/get_states.md)
    - [Create State](routes/state/create_state.md)
    - [Update State](routes/state/update_state.md)
    - [Delete State](routes/state/delete_state.md)
  - **Tag**: Manage tags for tasks.
    - [Get Tags](routes/tag/get_tags.md)
    - [Create Tag](routes/tag/create_tag.md)
    - [Update Tag](routes/tag/update_tag.md)
    - [Delete Tag](routes/tag/delete_tag.md)
- **[Database](database/index.md)**: Overview of MinBan's SQLite database schema and relationships.


## Technologies
The backend leverages:

- **Go**: Efficient and high-performance backend language.
- **Gin Framework**: A fast and simple web framework for building APIs.
- **SQLite**: A lightweight database system for storing data locally.
- **JWT Authentication**: For securing user access.

## Contribution

MinBan is an open-source project. Contributions are always welcome! Feel free to submit issues or pull requests on the [GitHub repository](#).