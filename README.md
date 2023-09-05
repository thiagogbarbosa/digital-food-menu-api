# Digital Food Menu API

A digital food menu API that provides real-time menu information for a restaurant. This API allows users to view the restaurant's menu items based on the current time, specific weekdays, or get the complete menu regardless of the day.

## Table of Contents

- [Features](#features)
- [Usage](#usage)
- [Packages Used](#packages-used)
- [Future Features](#future-features)
- [Installation](#installation)
- [Contributing](#contributing)
  
## Features

- Retrieve the menu items available for the current time.
- Get the complete menu, irrespective of the current time.
- Filter menu items based on specific weekdays.
- Returns menu data in JSON format.
- Lightweight and easy-to-use API.

## Usage

### Get Today's Menu

To retrieve the menu items available for the current time: `GET /menutoday`

### Get Complete Menu

To retrieve the complete menu: `GET /menu`

### Get Menu by Day

To retrieve menu items available on a specific day: `GET /menu/{available}`
Example: If I want the Monday menu -> `GET /menu/Monday`

## Packages Used

I plan to enhance this API with the following features:

- **Database Integration**: Implement a database(e.g., PostgreSQL, MySQL) to store menu data for better scalability and data management.
- **Authentication System**: Develop an authentication system to secure API access and allow restaurant staff to manage menu items securely
- **Unit Tests**: Create comprehensive unit tests to ensure the reliability and robustness of API functions.

## Installation

To run this API locally, follow these steps:

- Clone the repository: `git clone https://github.com/thiagogbarbosa/digital-food-menu-api.git`
- Navigate to the project directory: `cd digital-food-menu-api`
- Start the API server: `go run main.go`

The API server will start on port 8000 by default.

## Contributing

Constributions are welcome! Fell free to open issues or pull requests to help improve this project. If you have any ideas, bug reports, or feature requests, please shase them.


