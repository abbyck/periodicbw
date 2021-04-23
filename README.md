# PeriodicBW

A script made to run official speedtest.net binary periodically and store the results in a CSV file 

## Installation

Get the official speedtest binary for your architecture from [Speedtest CLI](https://www.speedtest.net/apps/cli)

```bash
git clone https://github.com/abbyck/periodicbw.git
```

## Usage

Change the cronjob duration in `main.go`
```bash
go build main.go
./main
```
Now speedtest results after the specified intervals will be saved the directory at which the compiled binary is run with the filename `<ISP_name>.csv`

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)