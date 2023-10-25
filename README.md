# License System
<p>Hello, i've been working with this project and will now put it out to the public.
I have made it for people that is interested in securing there FiveM scripts but you can also use it for other usecases than FiveM scripting by changing the</p> 
```lua
UsingFiveM```. <p>in the config.json.</p>

<p>It's not something that you should use for primarily safety but i will work on getting it to a higher security level. I have begun a website for it where you can Delete, Add licenses and so on, this is something that will need to running private since im not working on making the login more safe.</p>

<p>The project is meant to be runned on a host of some sort if personally use a VPS. Its made in GOLang so you will need to have some GO knowledge for setting it up to your preferences.</p>

## Installation

<p>Download the newest version on <github.link> to either your pc or a server, then if you want to change some of the code you will need to have installed GO on your pc and download the source code and from there build it! If thats not the case then dont worry.</p>

<p>1. Firstly you will need to import the db.sql inside of your database.</p>

<p>2. Setup your config.json to your preferences make sure that the DB url is setted up right: 
localhost: `dbname:password@localhost/dbname` remote db: `name:password@tcp(ipadress:port)/dbname`</p>

<p>3. Run the installed license.exe and see it working.</p>

<p>If you run it on localhost you will need to use this url `localhost:port/api/{resourcename}/{license}/{ipadresss}` dont mind the {resourcename} if not using FiveM in the config.json then it will just be `localhost:port/api/{license}/{ipadresss}`</p>

<p>If you need any help contact me on discord via my Discord server: https://discord.gg/XW9WGTrrmJ</p>

## Usage
<p>If you want to add new licenses then you will need to add them manually inside the Database for now i'm working on a website for it!</p>

<h8>Examples (Server sided)</h8>
```lua

AddEventHandler('onResourceStart', function(resourceName)
    if (GetCurrentResourceName() ~= resourceName) then
      return
    end

  local CurrentResouceName = GetCurrentResourceName()

    PerformHttpRequest("localhost:8080/api/".. CurrentResouceName .."/LQC1h2WQwDBxEm6QlDhM/123.456.789", function (errorCode, resultData, resultHeaders, errorData)
        Result = resultData
        if Result == "false" then 
            print("Invalid licensekey contact the script owner to fix this issue!")
            -- Then you could add something else like deleting the script or something
        else    
            print(resourceName .. ' has successfully loaded.')
        end
      end)
    end)

```

<p>Feel free to contribute with your own examples and ideas.</p>

## How it works
<p>The way it works is that you will implement it inside of your script (Examples comming soon), and then it will return back with either a "invalid" or "valid" string.
You can then use that string to do what ever you want.</p> 

<p>But its basicaly a API you call it like this example (In this example we have "UsingFiveM": true inside the config.json) localhost:8080/api/fivem-bungeejump/LQC1h2WQwDBxEm6QlDhM/123.456.789 
Okay so lets split it up</p> 
<p>• we first have the ip which in our case is localhost because i run it local,</p> 
<p>• and then we have a PORT which can be defined inside the config.json,</p>
<p>• then we have the resourcename, and after that we have our licensekey,</p>
<p>• and lastly the ipadress.</p>

<p>What then happens is that we make a call to the server running the license system where we send the resourcename the licensekey and the ipadress to the program. It then starts of checking if the licensekey exists if it does then we move on to checking the ipadress if the ipadress that is setted up to the licensekey in the database is the same as the API caller then we move on to checking if the resourcename sent is the same as the resourcename associated to the licensekey, if all of that succeeds then we return "valid" else it will just be a "invalid".
</p>

## Contributing

<p>I would love to see contributions to my project.</p>

<p>Please make sure to update tests as appropriate.</p>
