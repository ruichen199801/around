# Around

Project Date: May 2021

## Description

<img src="./images/demo.png" alt="demo" width="500"><br>

Around is a fullstack web app for sharing photos, videos, and posts. It uses React for the frontend and Go for the backend, and it supports fuzzy search with Elasticsearch.

## Architecture

<img src="./images/architecture.png" alt="architecture" width="500"><br>

### Frontend

Around features user-friendly webpages to create, browse, and search posts with React. Authentication is improved with token-based registration/login/logout and server-side authentication with JWT. The app is deployed on AWS Amplify.

### Backend

Around uses a scalable web service in Go for managing user posts and account information. Elasticsearch provides advanced search features for posts, and the app is deployed on Google App Engine for scalability.