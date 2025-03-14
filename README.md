# Go Expert Course - FullCycle

## Project Overview
This project is part of the postgraduate course in Go at FullCycle. The main objective is to apply the knowledge acquired about HTTP web servers, contexts, databases, and file manipulation with Go.  

## Challenge Summary

### client.go
- Makes an HTTP request to server.go to get the dollar exchange rate.
- Receives the current exchange rate value and saves it in a file named cotacao.txt.
- Uses the context package to handle timeouts and logs errors if the execution time is insufficient.

### server.go
- Consumes an external API to get the Dollar to Real exchange rate.
- Returns the exchange rate in JSON format to the client.
- Uses the context package to handle timeouts for API calls and database operations, logging errors if the execution time is insufficient.
- Registers each received exchange rate in an SQLite database.