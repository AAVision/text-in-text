# Text-in-Text :star:

**Text-in-text is a CLI tool built using GO. The idea behind that tool is to hide text inside text and use plaintext steganography.**


<p align="center">
  <img src="https://github.com/AAVision/text-in-text/blob/665b7d06f2750195109c8e44c23e62f6143ef77a/text-in-text.png" />
</p>


## Used Technologies :hammer_and_wrench:
- Go 1.22.4
- Cobra 1.8.1

## Usage :rocket:

The tool can encode and decode the texts, and it takes the following attributes:

- Encoding:
```bash
go run .\main.go encode --text "This is a cover text." --secret "Secret" --password "aavision"
```
```bash
Usage:
  text-in-text encode [flags]

Flags:
  -h, --help            help for encode
      --secret string   A secret to be hidden!

Global Flags:
      --password string   A password to protect your text!
      --text string       An encoded text to extract secret from it!
```

**The encoding will generate a file with a timestamp name, and it will contain the cover text encoded.**

- Decoding:

```bash
go run .\main.go decode --path "1234111.txt" --password "aavision"
```

```bash
Usage:
  text-in-text decode [flags]

Flags:
  -h, --help          help for decode
      --path string   Path of file

Global Flags:
      --password string   A password to protect your text!
      --text string       An encoded text to extract secret from it!
```

## Build :whale:

```bash
go build
```

## LICENSE :balance_scale:

This project is licensed under the MIT License. See the [LICENSE](https://github.com/AAVision/text-to-text/blob/main/LICENSE) file for details.
