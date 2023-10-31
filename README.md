# License System

Welcome to the License System project! This system is designed to help secure your FiveM scripts, but it can also be adapted for various other use cases by changing the `"UsingFiveM": <bool>` setting in the `config.json`. While it's not intended to be the ultimate security solution, it is a very good free solution. 
<p>Additionally, I have started work on a website where you can manage licenses, including adding and deleting them. Please note that the login system for this website should be strengthened if you plan to use it publicly.</p>

This project is designed to run on a server or host, and personally, I prefer to use a VPS. It's implemented in Golang, so some knowledge of Go is necessary to set it up according to your requirements.

## Installation

1. Download the latest version from [GitHub](https://github.com/ledepede1/golicense/releases/latest) to either your local machine or a server. If you need to make code changes, make sure you have Go installed on your PC and download the source code. If not, no worries.

2. Import the `db.sql` file into your database.

3. Configure the `config.json` to your preferences. Make sure the DB URL is correctly set,
   - whether it's a local database (e.g., `dbname:password@localhost/dbname`) 
   - or a remote one (e.g., `name:password@tcp(ipaddress:port)/dbname`).

5. Run the `license.exe` to see the system in action.

If you're running it on localhost, use this URL format: `localhost:port/api/{resourcename}/{license}`. If you're not using FiveM, it will be: `localhost:port/api/{license}`. If you need assistance, contact me on Discord via my [Discord server](https://discord.gg/XW9WGTrrmJ).

## Usage

Currently, if you want to add new licenses, you'll need to do it manually in the database. I'm actively working on a website to simplify this process.

### Examples (Server Sided)

```lua
AddEventHandler('onResourceStart', function(resourceName)
    if (GetCurrentResourceName() ~= resourceName) then
      return
    end

    local CurrentResourceName = GetCurrentResourceName()

    PerformHttpRequest("localhost:8080/api/".. CurrentResourceName .."/LQC1h2WQwDBxEm6QlDhM/123.456.789", function (errorCode, resultData, resultHeaders, errorData)
        Result = resultData
        if Result == "false" then 
            print("Invalid license key. Contact the script owner to resolve this issue.")
            -- You could implement additional actions like disabling the script.
        else    
            print(resourceName .. ' has successfully loaded.')
        end
      end)
end)
```

Feel free to contribute with your own examples and ideas.

## How It Works

The system is implemented by integrating it into your script (examples are coming soon). It returns either an "invalid" or "valid" string, which you can use in various ways.

This system essentially works as an API. If you're using FiveM (set 'UsingFiveM' to true in the `config.json`), you make a request like this: `localhost:8080/api/fivem-bungeejump/LQC1h2WQwDBxEm6QlDhM`. Let's break it down:

- The IP is in my case set to localhost, as I run it locally.
- The port can be defined in the `config.json`.
- The resource name is next in the URL, followed by the license key.

The system then checks whether the license key exists, followed by an IP address check. If the caller IP matches the one associated with the license key in the database, and the resource name matches the one sent, it returns "valid"; otherwise, it returns "invalid".

## Contributing

I would greatly appreciate contributions to this project. Please make sure to update tests as needed and improve the project as you see fit. Together, we can make it even better.
