package main

func main() {
	println(TranslateCodec("QasdfaiDSIAU", EncoderDictionary))
	println(TranslateCodec("Qasdfa0245ODS", DecoderDictionary))
	println(TranslatePigLatin("das fdsg fsdgh asdf"))
	println(TranslatePigLatin("ds ghfdgfj asdfdg"))
}
