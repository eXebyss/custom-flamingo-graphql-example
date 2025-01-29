# Custom Flamingo GraphQL Example

This project extends the original Flamingo GraphQL example with additional functionality and custom packages.

## Technology Stack

### Backend
- Golang
- Flamingo Framework
- Flamingo Ecommerce
- Flamingo GraphQL

### Frontend
- HTML templates (served from backend)
- TailwindCSS
- Alpine.js

## Key Differences from Original Example
- Uses custom package implementations instead of default Flamingo packages
- Added new mutations for todo management:
    - Delete todo items
    - Edit existing todo items

## Setup and Running

Clone the repository:
```bash
git clone git@github.com:eXebyss/custom-flamingo-graphql-example.git
сd custom-flamingo-graphql-example
```

Generate GraphQL code:
```bash
go generate .
```

## Available Endpoints
- GraphQL Console: http://localhost:3322/graphql-console
- Frontend: http://localhost:3322/

## GraphQL Operations
- Query todos
- Add new todos
- Mark todos as done/undone
- Delete todos
- Edit todo content

## Development
When adding new fields or structs, regenerate the GraphQL code:
```bash
go generate .
```

## Tutorial
https://de.slideshare.net/i-love-flamingo/graphql-with-flamingo

## Original Example
https://github.com/i-love-flamingo/graphql.git
