# blue-brick

Bluebricks CLI allows users to log in using a browser-based authentication system.


## Installation

1. Clone the repo:
   ```sh
   git clone https://github.com/tanecar/bluebrick.git
   cd bluebrick

## Build executable

go build -o bluebrick

## Usage

./bluebrick login  - Opens the browser for authentication and stores the token.
./bluebrick status - Shows if the user is logged in.
./bluebrick logout - Logouts user