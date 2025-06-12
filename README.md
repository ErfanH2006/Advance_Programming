# Personal Finance Management CLI (in Go)

## 🔍 Project Overview
This is a command-line tool written in Golang that helps users manage personal finances by tracking incomes, expenses, categories, and generating reports from data stored in JSON files.

## 🧩 Features

### 👤 User Management
- Add user with username and email
- List all users
- Switch active user

### 💰 Transaction Management
- Add income (with amount, source, date, description)
- Add expense (with amount, category, date, description)
- View all transactions
- Delete any transaction by ID

### 🗂️ Category Management
- Default categories (Food, Transport, Bills, Shopping, Entertainment, Other)
- Add custom categories per user

### 📊 Reports
- **Monthly Summary**: Shows income, expense and balance per month
- **Category Report**: Shows total expenses per category
- **Daily Balance**: Shows balance for each transaction date

## 🛠 Technical Info

- **Language:** Go (Golang)
- **Data Storage:** Local JSON files (no database)
- **Architecture:** Follows OOP principles: base class `Transaction`, derived classes `Income` and `Expense`

## 🏗 Project Structure
