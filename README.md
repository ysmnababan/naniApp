# ChatApp (Under Construction 🚧)

This project is currently under development. ChatApp is a simple messaging platform where users can chat in real-time. It allows users to create accounts, send messages, and interact with each other. The backend is developed in Golang, utilizing WebSockets for real-time communication. 

## Current Status

- [x] Initial project setup
- [x] User registration & authentication
- [ ] WebSocket-based real-time chat
- [ ] Message storage and history
- [ ] API documentation



## Features

- User account creation and authentication
- Real-time messaging with WebSockets
- Secure message storage
- RESTful API for account management
- Connection pooling for optimized performance
- Redis for caching and message management
- Elasticsearch for indexing and searching chat history
- Support for microservices architecture

## Tech Stack (Planned)

- **Backend:** Golang
- **WebSockets:** For real-time chat
- **Database:** PostgreSQL for storing user data and messages
- **Caching:** Redis for efficient message management
- **Search:** Elasticsearch for fast searching of messages
- **Frontend:** React (connecting WebSocket to display messages)
- **Deployment:** Docker, Kubernetes (optional for production)

## Getting Started

This project is still in its early stages. Check back later for updates!

### Prerequisites

- Go 1.18+
- PostgreSQL / MySQL
- Redis
- Elasticsearch
- Docker (for running the app in containers)


### API Gateway

The **API Gateway** serves as the single entry point for clients and routes requests to the appropriate microservice based on the endpoint accessed:

#### User Service Endpoints

- `POST /api/users/register`: Create a new user account.
- `POST /api/users/login`: Authenticate a user and log them in.
- `GET /api/user/`: Fetch the user profile.
- `PUT /api/user/`: Edit user profile.

#### Message Service Endpoints

The **Message Service** interacts with the **User Service** to retrieve user data and the **Notification Service** to send push notifications when messages are sent.

- `POST /api/messages/send`: Store the message in the database.
- `GET /api/messages/{conversationId}`: Fetch all messages from a specific conversation (private chat or group).
- `PUT /api/messages/unsent/{messageId}`: Unsend a message.
- `GET /api/messages/member`: Fetch the list of all members in a group.
- `GET /api/messages/read/{messageId}`: Fetch the list of all members who have read a specific message.
- `POST /api/participant/{conversationId}`: Add a new member to a group.

### WebSocket Usage
ws://your-server-url/ws
To connect to the WebSocket, use the following endpoint:

- Send a message: `{ "type": "message", "data": "Hello, world!" }`
- Receive a message: `{ "user": "John", "message": "Hi!" }`

### Database Structure

- **Users**: Stores user information.
- **Conversations**: Represents both private and group conversations.
- **Participants**: Links users to conversations.
- **Messages**: Stores message content and metadata.
- **Message Status**: Tracks the read/delivered status of each message for each participant.
> **For detailed ERD and decision-making documentation**, please refer to the [docs](./docs) folder.

### Contribution

Feel free to fork this repository and contribute by submitting pull requests. Ensure your code follows the project's coding guidelines.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.