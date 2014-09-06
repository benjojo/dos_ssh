package main

import (
	"encoding/hex"
)

func VESAtoVT100(code byte) string {
	// Super hax lol
	template := "\x1B["
	out := hex.EncodeToString([]byte{code})
	if out[:1] == "0" {
		template = template + "40"
	} else if out[:1] == "1" {
		template = template + "44"
	} else if out[:1] == "2" {
		template = template + "42"
	} else if out[:1] == "3" {
		template = template + "46"
	} else if out[:1] == "4" {
		template = template + "41"
	} else if out[:1] == "5" {
		template = template + "45"
	} else if out[:1] == "6" {
		template = template + "43"
	} else if out[:1] == "7" {
		template = template + "47"
	} else if out[:1] == "8" {
		template = template + "40"
	} else if out[:1] == "9" {
		template = template + "44"
	} else if out[:1] == "A" {
		template = template + "42"
	} else if out[:1] == "B" {
		template = template + "46"
	} else if out[:1] == "C" {
		template = template + "41"
	} else if out[:1] == "D" {
		template = template + "45"
	} else if out[:1] == "E" {
		template = template + "43"
	}

	if out[1:] == "0" {
		template = template + ";30"
	} else if out[1:] == "1" {
		template = template + ";34"
	} else if out[1:] == "2" {
		template = template + ";32"
	} else if out[1:] == "3" {
		template = template + ";36"
	} else if out[1:] == "4" {
		template = template + ";31"
	} else if out[1:] == "5" {
		template = template + ";35"
	} else if out[1:] == "6" {
		template = template + ";33"
	} else if out[1:] == "7" {
		template = template + ";37"
	} else if out[1:] == "8" {
		template = template + ";30"
	} else if out[1:] == "9" {
		template = template + ";34"
	} else if out[1:] == "A" {
		template = template + ";32"
	} else if out[1:] == "B" {
		template = template + ";36"
	} else if out[1:] == "C" {
		template = template + ";31"
	} else if out[1:] == "D" {
		template = template + ";35"
	} else if out[1:] == "E" {
		template = template + ";33"
	}
	template = template + "m"
	return template
	/*
			Set Display Attributes

		Set Attribute Mode	<ESC>[{attr1};...;{attrn}m
		Sets multiple display attribute settings. The following lists standard attributes:
		0	Reset all attributes
		1	Bright
		2	Dim
		4	Underscore
		5	Blink
		7	Reverse
		8	Hidden

			Foreground Colours
		30	Black
		31	Red
		32	Green
		33	Yellow
		34	Blue
		35	Magenta
		36	Cyan
		37	White

			Background Colours
		40	Black
		41	Red
		42	Green
		43	Yellow
		44	Blue
		45	Magenta
		46	Cyan
		47	White
	*/
}

func CorrectBadChars(in byte) string {

	data := []string{" ", "☺", "☻", "♥", "♦", "♣", "♠", "•", "◘",
		"○", "◙", "♂", "♀", "♪", "♬", "☼", "►", "◄", "↕", "‼", "¶", "§",
		"▬", "↨", "↑", "↓", "→", "←", "∟", "↔", "▲", "▼", " ", "!", "\"",
		"#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":", ";", "<",
		"=", ">", "?", "@", "A", "B", "C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z", "[", "\\", "]", "^", "_", "`", "a", "b", "c",
		"d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "{", "|", "}",
		"~", "⌂", "Ç", "ü", "é", "â", "ä", "à", "å", "ç", "ê", "ë", "è",
		"ï", "î", "ì", "Ä", "Å", "É", "æ", "Æ", "ô", "ö", "ò", "û", "ù",
		"ÿ", "Ö", "Ü", "¢", "£", "¥", "₧", "ƒ", "á", "í", "ó", "ú", "ñ",
		"Ñ", "ª", "º", "¿", "⌐", "¬", "½", "¼", "¡", "«", "»", "░", "▒",
		"▓", "│", "┤", "╡", "╢", "╖", "╕", "╣", "║", "╗", "╝", "╜", "╛",
		"┐", "└", "┴", "┬", "├", "─", "┼", "╞", "╟", "╚", "╔", "╩", "╦",
		"╠", "═", "╬", "╧", "╨", "╤", "╥", "╙", "╘", "╒", "╓", "╫", "╪",
		"┘", "┌", "█", "▄", "▌", "▐", "▀", "α", "ß", "Γ", "π", "Σ", "σ",
		"µ", "τ", "Φ", "Θ", "Ω", "δ", "∞", "φ", "ε", "∩", "≡", "±", "≥",
		"≤", "⌠", "⌡", "÷", "≈", "°", "∙", "·", "√", "ⁿ", "²", "■", "\n", " "}

	return data[uint8(in)]
}
