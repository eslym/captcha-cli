# Captcha CLI

A very simple CLI tool to generate captcha image and output it to `STDOUT`.

## Usage

```bash
captcha -data-url -width 200 -height 100 123456
```
> `captcha [options] <captcha-text>`
>
> where captcha-text must be a number.
> 
> **Options:**
> - `-data-url` output captcha image as data url, otherwise output as raw png image
> - `-width` captcha image width
> - `-height` captcha image height
