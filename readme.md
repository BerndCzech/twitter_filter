### Usage

Assuming `.env` looks like: 

TOKEN=YourBEARERTokenWithNoQuotes


```shell
grep export $(grep TOKEN .env);
go build --o tFilter && ./tFilter
```