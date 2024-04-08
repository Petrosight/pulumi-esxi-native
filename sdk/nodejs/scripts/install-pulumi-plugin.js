"use strict";

var childProcess = require("child_process");

/* 0 => type, 1 => name, 2 => version */
var args = process.argv.slice(2);

if (args.indexOf("${VERSION}") !== -1) {
    process.exit(0);
}
if (args.length !== 3) {
    console.error('Incorrect number of arguments, expected 3 received ' + args.length)
    process.exit(1)
}

const wantedVersion = args[2].match(/^v?(\d+(?:\.\d+)+(?:-[a-zA-Z0-9]+)?).*/)[1]
try {
    const output = childProcess.spawnSync("pulumi", ["plugin", "ls", "-j"], {
        stdio: ["ignore", "pipe", "inherit"]
    }).stdout.toString()
    const check = JSON.parse(output)
    if (Array.isArray(check)) {
        check.forEach(entry => {
            if (
                entry.kind === args[0] &&
                entry.name === args[1] &&
                entry.version === wantedVersion
            ) {
                console.log("Found plugin already installed")
                process.exit(0)
            }
        })
    }
} catch (e) {
    /* Do nothing */
    console.error('Exception: ' + e)
}

var res = childProcess.spawnSync("pulumi", ["plugin", "install", "--server", "github://api.github.com/petrosight/pulumi-esxi-native"].concat(args), {
    stdio: ["ignore", "inherit", "inherit"]
});

if (res.error && res.error.code === "ENOENT") {
    console.error("\nThere was an error installing the resource provider plugin. " +
        "It looks like `pulumi` is not installed on your system. " +
        "Please visit https://pulumi.com/ to install the Pulumi CLI.\n" +
        "You may try manually installing the plugin by running " +
        "`pulumi plugin install " + args.join(" ") + "`");
} else if (res.error || res.status !== 0) {
    console.error("\nThere was an error installing the resource provider plugin. " +
        "You may try to manually installing the plugin by running " +
        "`pulumi plugin install " + args.join(" ") + "`");
}

process.exit(0);
