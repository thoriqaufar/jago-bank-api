
# Inventory Management App

## Tech Stack
- PHP Laravel
- HTML
- CSS
- React
- MySQL



## Run Locally

Clone the project

```bash
git clone git@github.com:thoriqaufar/inventory-management.git inventory-management
```

Go to the project directory

```bash
cd inventory-management
```

Install backend dependencies

```bash
composer install
```

Setup environment

```bash
cp .env.example .env
```

Update database information in .env

```bash
DB_CONNECTION=your_db
DB_HOST=your_db_host
DB_PORT=your_db_port
DB_DATABASE=your_db_name
DB_USERNAME=your_db_username
DB_PASSWORD=your_db_password
```

Run migration and seeder

```bash
php artisan migrate
php artisan db:seed
```

Start the backend server

```bash
php artisan serve
```

Install frontend dependencies

```bash
cd frontend
npm i
```

Start the frontend server

```bash
npm run dev
```
