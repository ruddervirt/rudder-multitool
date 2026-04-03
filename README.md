rudder-multitool
=================

Minimal scaffold for a cross-platform multitool CLI.

Build:

```bash
GOOS=windows GOARCH=amd64 go build -o rudder-multitool.exe
GOOS=linux GOARCH=amd64 go build -o rudder-multitool-linux
GOOS=darwin GOARCH=amd64 go build -o rudder-multitool-mac
```

Usage:

```bash
./rudder-multitool version
```


This project is licensed under the GNU General Public License v3.0 (GPL-3.0).

See the LICENSE file for details or visit https://www.gnu.org/licenses/gpl-3.0.html for the full license text.