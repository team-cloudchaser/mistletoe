# Winter Mistletoe
🔗 A simple CLI dialer. Supports Linux, Windows and Android (Termux).

## Usage
- `./update`: Updater for fetching all configurations from the master pointer. (WIP)
- `./load`: Loader of the actual dialer. Help availble with `./load help`.
- `./conf/<tId>.json`: The matched dialer config file.
- `./conf/default.json`: The default dialer config when no matched dialer configs are present.
- `./data/<tId>.tsv`: The matched dialer target list.
- `./data/default.tsv`: The default dialer target list when none matched.
- `./data/<tId>.json`: The matched dialer template file.

## File schema
### Dialer config
```hjson
{
	"run": ["binary", "arg1", "%s", "arg2"],
	"windowsPrefix": "example/windows-path",
	"linuxPrefix": "example/linux-path"
}
```
- `run`: The command invoked after config generation is finished. The first string contains the executable file to run, while the rest denotes arguments, with `%s` replaced with the name of the generated config file.
- `windowsPrefix`: (optional) The executable path prefix when on Windows. Delimited with forward slashes `/`.
- `linuxPrefix`: (optional) The executable path prefix when on Linux or Android.

### Target list
```tsv
id	field_1	field_2
pun	chips	cheese
```
A TSV file with the first line denoting fields. The `id` field must always be present and be the first field.

Values of the rest of the fields will replace all occurances of the capitalized `__FIELD_NAME__` placeholders, like when `pun` is selected as the target, all occurances of `__FIELD_1__` will be replaced with `chips`, and `__FIELD_2__` with `cheese`.

### Templates
Any valid JSON config file with replaceable placeholder fields is viable.

For writing valid Xray config files, head to [`xtls.github.io`](https://xtls.github.io/en/config/).