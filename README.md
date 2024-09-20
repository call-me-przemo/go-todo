# Go Todo

![hero](https://res.cloudinary.com/ichtrojan/image/upload/v1574958373/Screenshot_2019-11-28_at_17.22.25_gyegdr.png)

## Introduction

A simple todolist application written in Go

## Requirements

- Docker installed
- Go installed

## Installation

- Clone this repo

```bash
git clone https://github.com/call-me-przemo/go-todo.git
```

- Change Directory

```bash
cd go-todo
```

- Initiate `.env` file

```bash
cp .env.example .env
```

- Modify `.env` file with your correct database credentials and desired Port

## Usage

Run database

```bash
docker compose up -d
```

To run this application, execute:

```bash
go run .
```

You should be able to access this application at `http://0.0.0.0:3000`

**NOTE**

If you modified the port in the `.env` file, you should access the application for the port you set
