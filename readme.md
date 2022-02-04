<!-- ABOUT THE PROJECT -->

## About The Project

GraphQL API for Gaskeuun, an event planning App created for the purpose of study.

Building the project with layered architecture, and clean code approach for the structure, with the intention of simplicity when the app is scaling up and ease of maintenance

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

This project structure is built using

- [GraphQL]
- [Golang]
- [Mysql]
- [Labstack/Echo]

<p align="right">(<a href="#top">back to top</a>)</p>

### Features

- USERS CRUD
- EVENTS CRUD
- PARTICIPANT CRUD
- COMMENTS CR

### Folder Structure

```
├── config/                        # fill
├
├── delivery/                      # fill
├──── controllers/                 # fill
├──── middlewares/                 # fill
├──── router/                      # fill
├
├── entities/                      # fill
├──── model/                       # fill
├
├── repository/                    # fill
├──── auth                         # fill
├──── user                         # fill
├
├── util/                          # fill
├──── graph/                       # fill
├──────── generated/               # fill

```

<!-- GETTING STARTED -->

## Getting Started

To start project, just clone this repo

### Installation

# Clone the repo

```bash
git clone https://github.com/HamzahAA15/Event-Planning-App.git
```

# How To Edit

Step for edit :

```bash
1. buka folder util/graph
2. pilih file schema.graphqls
3. edit sesuai yang diinginkan
4. generate ulang
```

# GraphQL command

Generate Ulang :

```bash
go run github.com/99designs/gqlgen generate
```

<p align="right">(<a href="#top">back to top</a>)</p>
