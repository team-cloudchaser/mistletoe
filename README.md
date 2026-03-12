# Winter Mistletoe
🔗 Dialer toolchain in Bash for POSIX. Supports Linux and Android (Termux).

## Usage
- `./update`: Updates all configurations from the pointer.
- `./load`: The loader of the actual dialer.

## Configuration
- `conf/pointer.txt`: Stores the URL for the full list of files with file names to download. URLs can contain an `${auth}` variable to bring in the authentication string.
- `conf/auth.txt`: Stores the authentication string supplied to `pointer.txt`.
- `conf/run.sh`: Contains the actual script running the dialer. `$1` contains the JSON config template, `$2` contains the target specifier, and `$MISTLETOE_LINE` contains the unparsed tab-delimited line.
- `data/list.tsv`: Downloaded from `pointer.txt`. Where to fetch the actual configuration templates and dynamic params. URLs can contain an `${auth}` variable to bring in the authentication string. The file should end with a line feed.
- `data/remote.tsv`: Downloaded from `pointer.txt`. The actual targets to connect to from the specifier. The first field should always be the target specifier, and the file should end with a line feed.