<div align="center">
<img src="logo.png" style="width: 30%"><br>
A CLI web-reconnaissance tool written in Go
<br>
<br>
<img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/t/Sanyam-Asthana/SammyGo">
<img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Sanyam-Asthana/SammyGo">
<img alt="GitHub Created At" src="https://img.shields.io/github/created-at/Sanyam-Asthana/SammyGo">

<br>

<img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/Sanyam-Asthana/SammyGo">
<a href="https://github.com/Sanyam-Asthana/SammyGo/actions/workflows/go.yml"><img alt="GitHub Actions" src="https://github.com/Sanyam-Asthana/SammyGo/actions/workflows/go.yml/badge.svg"></a>

</div>

# Introduction
SammyGo is a web-reconnaissance tool with functionality like requests viewing, directory busting, web crawling etc. in one package.

It is a complete rewrite of its previous iteration, [Sammy](https://github.com/Sanyam-Asthana/sammy-web).

# Instructions

The project was created and tested using `go 1.25.5`

## Compiling
1. Clone the repository
2. 	Once in the root directory, run the following:

	```bash
	cd cmd/sammygo/
	go build -o ../../build/sammygo .
	```
3. The compiled binary will be located in `[repository root]/build/`

## Usage

### Checking status
Run `sammygo [URL]`

```
❯ ./sammygo http://scanme.nmap.org
 ____                                   ____
/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___
\___ \ / _` | '_ ` _ \| '_ ` _ \| | | | |  _ / _ \
 ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |
|____/ \__,_|_| |_| |_|_| |_| |_|\__, |\____|\___/
                                 |___/

Initiated SammyGo on http://scanme.nmap.org

--------------------
Status Code: 200 OK
--------------------
```

### Getting response header
Run `sammygo --head [URL]`

```
❯ ./sammygo --head http://scanme.nmap.org
 ____                                   ____
/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___
\___ \ / _` | '_ ` _ \| '_ ` _ \| | | | |  _ / _ \
 ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |
|____/ \__,_|_| |_| |_|_| |_| |_|\__, |\____|\___/
                                 |___/

Initiated SammyGo on http://scanme.nmap.org

--------------------
Status Code: 200 OK
--------------------

----------HEADER----------
Vary: Accept-Encoding
Content-Type: text/html
Date: Thu, 26 Feb 2026 14:23:02 GMT
Server: Apache/2.4.7 (Ubuntu)
Accept-Ranges: bytes
--------------------------
```

### Getting response text
Run `sammygo --text [URL]`

```
❯ ./sammygo --text http://scanme.nmap.org
 ____                                   ____
/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___
\___ \ / _` | '_ ` _ \| '_ ` _ \| | | | |  _ / _ \
 ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |
|____/ \__,_|_| |_| |_|_| |_| |_|\__, |\____|\___/
                                 |___/

Initiated SammyGo on http://scanme.nmap.org

--------------------
Status Code: 200 OK
--------------------

-------------BODY-------------
<!DOCTYPE html>
<html lang="en">
<head>
<title>Go ahead and ScanMe!</title>
<meta name="viewport" content="width=device-width,initial-scale=1">
<meta name="theme-color" content="#2A0D45">

...(truncated for readability)

</body>
</html>


------------------------------
```

### Directory Busting
#### Default mode
`Run sammygo --bust [URL]`

```
❯ ./sammygo --bust http://scanme.nmap.org
 ____                                   ____
/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___
\___ \ / _` | '_ ` _ \| '_ ` _ \| | | | |  _ / _ \
 ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |
|____/ \__,_|_| |_| |_|_| |_| |_|\__, |\____|\___/
                                 |___/

Initiated SammyGo on http://scanme.nmap.org

--------------------
Status Code: 200 OK
--------------------

Initiating buster with 20 workers...
[+] Found: http://scanme.nmap.org/.%2e/%2e%2e/%2e%2e/%2e%2e/etc/passwd (400)
[+] Found: http://scanme.nmap.org/.htaccess (403)
[+] Found: http://scanme.nmap.org/.htaccess.bak (403)
[+] Found: http://scanme.nmap.org/.htaccess-dev (403)

...(truncated for readability)

[+] Found: http://scanme.nmap.org/images/ (200)
[+] Found: http://scanme.nmap.org/images (200)
[+] Found: http://scanme.nmap.org/index (200)
[+] Found: http://scanme.nmap.org/index.html (200)
[+] Found: http://scanme.nmap.org/shared (403)

Scan complete.
```

#### Filter by status

As is evident from previous example, most of the brute-forced paths are returning 403 (Forbidden) as a status code. If only the accessible paths are required, `--filter` flag may be used.

Run `sammygo --bust --filter=[Status Code] [URL]`

```
❯ ./sammygo --bust --filter=200 http://scanme.nmap.org
 ____                                   ____
/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___
\___ \ / _` | '_ ` _ \| '_ ` _ \| | | | |  _ / _ \
 ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |
|____/ \__,_|_| |_| |_|_| |_| |_|\__, |\____|\___/
                                 |___/

Initiated SammyGo on http://scanme.nmap.org

--------------------
Status Code: 200 OK
--------------------

Initiating buster with 20 workers...
[+] Found: http://scanme.nmap.org/images/ (200)
[+] Found: http://scanme.nmap.org/images (200)
[+] Found: http://scanme.nmap.org/index (200)
[+] Found: http://scanme.nmap.org/index.html (200)

Scan complete.
```

#### Number of workers
By default, the brute-forcer uses 20 workers (goroutines) at once. This number can be increased or decreased to change the speed of the buster. It should be kept in mind that high speeds can (and will almost always) cause rate-limiting.

`Run sammygo --bust --workers=[Number of workers] [URL]`

```
❯ ./sammygo --bust --workers=200 http://scanme.nmap.org
 ____                                   ____
/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___
\___ \ / _` | '_ ` _ \| '_ ` _ \| | | | |  _ / _ \
 ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |
|____/ \__,_|_| |_| |_|_| |_| |_|\__, |\____|\___/
                                 |___/

Initiated SammyGo on http://scanme.nmap.org

--------------------
Status Code: 200 OK
--------------------

Initiating buster with 200 workers...
[+] Found: http://scanme.nmap.org/.%2e/%2e%2e/%2e%2e/%2e%2e/etc/passwd (400)
[+] Found: http://scanme.nmap.org/.htaccess (403)
[+] Found: http://scanme.nmap.org/.htaccess.bak (403)
[+] Found: http://scanme.nmap.org/.htaccess-dev (403)

...(truncated for readability)

[+] Found: http://scanme.nmap.org/images/ (200)
[+] Found: http://scanme.nmap.org/images (200)
[+] Found: http://scanme.nmap.org/index (200)
[+] Found: http://scanme.nmap.org/index.html (200)
[+] Found: http://scanme.nmap.org/shared (403)

Scan complete.
```

#### Verbosity
The flag `--verbose` can be used to monitor the activity of the buster. The output of this feature is bugged at the moment.

Run `sammygo --bust --verbose [URL]`

```
❯ ./sammygo --bust --verbose http://scanme.nmap.org
 ____                                   ____
/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___
\___ \ / _` | '_ ` _ \| '_ ` _ \| | | | |  _ / _ \
 ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |
|____/ \__,_|_| |_| |_|_| |_| |_|\__, |\____|\___/
                                 |___/

Initiated SammyGo on http://scanme.nmap.org

--------------------
Status Code: 200 OK
--------------------

Initiating buster with 20 workers...
Checking: 41

[+] Found: http://scanme.nmap.org/.%2e/%2e%2e/%2e%2e/%2e%2e/etc/passwd (400)
Checking: 575

[+] Found: http://scanme.nmap.org/.htaccess (403)
Checking: 577

[+] Found: http://scanme.nmap.org/.htaccess.BAK (403)
Checking: 578
```

#### Custom wordlist
The repository comes with a default wordlist, but the wordlist can be customized using the `--wordlist` flag.

Run `sammygo --bust --wordlist=[Wordlist path] [URL]`

`[Wordlist path]` is the relative path to the program binary.

```
❯ ./sammygo --bust --wordlist=words.txt http://scanme.nmap.org
 ____                                   ____
/ ___|  __ _ _ __ ___  _ __ ___  _   _ / ___| ___
\___ \ / _` | '_ ` _ \| '_ ` _ \| | | | |  _ / _ \
 ___) | (_| | | | | | | | | | | | |_| | |_| | (_) |
|____/ \__,_|_| |_| |_|_| |_| |_|\__, |\____|\___/
                                 |___/

Initiated SammyGo on http://scanme.nmap.org

--------------------
Status Code: 200 OK
--------------------

Initiating buster with 20 workers...

Scan complete.
```

Needless to say, all of the flags can be used at once.

# Features Planned
All features of Sammy:

- Crawler
- Interactive shell
- Report generation

Plus,

- Cookies viewer
- Scraper

**Note:** SammyGo is intended for use only on systems you are allowed to test. Any use of web-reconnaissance tools like SammyGo on systems you are not authorized to test is unethical. The website used in the usage examples (http://scanme.nmap.org) allows testing of this kind.
