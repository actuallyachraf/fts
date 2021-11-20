package fts

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func compare(t *testing.T, expected, actual interface{}, msg ...string) {
	if expected != actual {
		t.Errorf("[%v] -- value differs. Expected [%v], actual [%v]", msg, expected, actual)
	}
}

func TestConsonant(t *testing.T) {
	word := []byte("TOY")
	compare(t, true, Consonant(word, 0), "T")  //T
	compare(t, false, Consonant(word, 1), "O") //O
	compare(t, true, Consonant(word, 2), "Y")  //Y
	word = []byte("SYZYGY")
	compare(t, true, Consonant(word, 0), "S")  //S
	compare(t, false, Consonant(word, 1), "Y") //Y
	compare(t, true, Consonant(word, 2), "Z")  //Z
	compare(t, false, Consonant(word, 3), "Y") //Y
	compare(t, true, Consonant(word, 4), "G")  //G
	compare(t, false, Consonant(word, 5), "Y") //Y
	word = []byte("yoke")
	compare(t, true, Consonant(word, 0), "YOKE")
}

func TestMeasure(t *testing.T) {
	compare(t, 0, Measure([]byte("TR")))
	compare(t, 0, Measure([]byte("EE")))
	compare(t, 0, Measure([]byte("TREE")))
	compare(t, 0, Measure([]byte("Y")))
	compare(t, 0, Measure([]byte("BY")))
	compare(t, 1, Measure([]byte("TROUBLE")))
	compare(t, 1, Measure([]byte("OATS")))
	compare(t, 1, Measure([]byte("TREES")))
	compare(t, 1, Measure([]byte("IVY")))
	compare(t, 2, Measure([]byte("TROUBLES")))
	compare(t, 2, Measure([]byte("PRIVATE")))
	compare(t, 2, Measure([]byte("OATEN")))
	compare(t, 2, Measure([]byte("ORRERY")))
}

func Test1A(t *testing.T) {
	compare(t, "caress", string(stepOneA([]byte("caresses"))))
	compare(t, "poni", string(stepOneA([]byte("ponies"))))
	compare(t, "ti", string(stepOneA([]byte("ties"))))
	compare(t, "caress", string(stepOneA([]byte("caress"))))
	compare(t, "cat", string(stepOneA([]byte("cats"))))
}

func Test1B(t *testing.T) {
	compare(t, "feed", string(stepOneB([]byte("feed"))))
	compare(t, "agree", string(stepOneB([]byte("agreed"))))
	compare(t, "plaster", string(stepOneB([]byte("plastered"))))
	compare(t, "bled", string(stepOneB([]byte("bled"))))
	compare(t, "motor", string(stepOneB([]byte("motoring"))))
	compare(t, "sing", string(stepOneB([]byte("sing"))))
	compare(t, "motor", string(stepOneB([]byte("motoring"))))
	compare(t, "conflate", string(stepOneB([]byte("conflated"))))
	compare(t, "trouble", string(stepOneB([]byte("troubled"))))
	compare(t, "size", string(stepOneB([]byte("sized"))))
	compare(t, "hop", string(stepOneB([]byte("hopping"))))
	compare(t, "tan", string(stepOneB([]byte("tanned"))))
	compare(t, "fail", string(stepOneB([]byte("failing"))))
	compare(t, "file", string(stepOneB([]byte("filing"))))
}

func Test1C(t *testing.T) {
	compare(t, "sky", string(stepOneC([]byte("sky"))))
	compare(t, "happi", string(stepOneC([]byte("happy"))))

}

func Test2(t *testing.T) {
	compare(t, "relate", string(stepTwo([]byte("relational"))))
	compare(t, "condition", string(stepTwo([]byte("conditional"))))
	compare(t, "rational", string(stepTwo([]byte("rational"))))
	compare(t, "valence", string(stepTwo([]byte("valenci"))))
	compare(t, "hesitance", string(stepTwo([]byte("hesitanci"))))
	compare(t, "digitize", string(stepTwo([]byte("digitizer"))))
	compare(t, "conformable", string(stepTwo([]byte("conformabli"))))
	compare(t, "radical", string(stepTwo([]byte("radicalli"))))
	compare(t, "different", string(stepTwo([]byte("differentli"))))
	compare(t, "vile", string(stepTwo([]byte("vileli"))))
	compare(t, "analogous", string(stepTwo([]byte("analogousli"))))
	compare(t, "vietnamize", string(stepTwo([]byte("vietnamization"))))
	compare(t, "predicate", string(stepTwo([]byte("predication"))))
	compare(t, "operate", string(stepTwo([]byte("operator"))))
	compare(t, "feudal", string(stepTwo([]byte("feudalism"))))
	compare(t, "decisive", string(stepTwo([]byte("decisiveness"))))
	compare(t, "hopeful", string(stepTwo([]byte("hopefulness"))))
	compare(t, "callous", string(stepTwo([]byte("callousness"))))
	compare(t, "formal", string(stepTwo([]byte("formaliti"))))
	compare(t, "sensitive", string(stepTwo([]byte("sensitiviti"))))
	compare(t, "sensible", string(stepTwo([]byte("sensibiliti"))))
}

func Test3(t *testing.T) {
	compare(t, "triplic", string(stepThree([]byte("triplicate"))))
	compare(t, "form", string(stepThree([]byte("formative"))))
	compare(t, "formal", string(stepThree([]byte("formalize"))))
	compare(t, "electric", string(stepThree([]byte("electriciti"))))
	compare(t, "electric", string(stepThree([]byte("electrical"))))
	compare(t, "hope", string(stepThree([]byte("hopeful"))))
	compare(t, "good", string(stepThree([]byte("goodness"))))
}

func Test4(t *testing.T) {
	compare(t, "reviv", string(stepFour([]byte("revival"))))
	compare(t, "allow", string(stepFour([]byte("allowance"))))
	compare(t, "infer", string(stepFour([]byte("inference"))))
	compare(t, "airlin", string(stepFour([]byte("airliner"))))
	compare(t, "gyroscop", string(stepFour([]byte("gyroscopic"))))
	compare(t, "adjust", string(stepFour([]byte("adjustable"))))
	compare(t, "defens", string(stepFour([]byte("defensible"))))
	compare(t, "irrit", string(stepFour([]byte("irritant"))))
	compare(t, "replac", string(stepFour([]byte("replacement"))))
	compare(t, "adjust", string(stepFour([]byte("adjustment"))))
	compare(t, "depend", string(stepFour([]byte("dependent"))))
	compare(t, "adopt", string(stepFour([]byte("adoption"))))
	compare(t, "homolog", string(stepFour([]byte("homologou"))))
	compare(t, "commun", string(stepFour([]byte("communism"))))
	compare(t, "activ", string(stepFour([]byte("activate"))))
	compare(t, "angular", string(stepFour([]byte("angulariti"))))
	compare(t, "homolog", string(stepFour([]byte("homologous"))))
	compare(t, "effect", string(stepFour([]byte("effective"))))
	compare(t, "bowdler", string(stepFour([]byte("bowdlerize"))))
}

func Test5A(t *testing.T) {
	compare(t, "probat", string(stepFiveA([]byte("probate"))))
	compare(t, "rate", string(stepFiveA([]byte("rate"))))
	compare(t, "ceas", string(stepFiveA([]byte("cease"))))
}

func Test5B(t *testing.T) {
	compare(t, "control", string(stepFiveB([]byte("controll"))))
	compare(t, "roll", string(stepFiveB([]byte("roll"))))
}

func TestVocal(t *testing.T) {
	f, err := os.Open("testdata/in.txt")
	if err != nil {
		t.Fatal("failed to open test file with error :", err)
	}
	in := bufio.NewReader(f)
	f, err = os.Open("testdata/out.txt")
	if err != nil {
		t.Fatal("failed to open test file with error :", err)
	}
	out := bufio.NewReader(f)
	for word, err := in.ReadSlice('\n'); err == nil; word, err = in.ReadSlice('\n') {
		stem, err := out.ReadSlice('\n')
		if err != nil {
			panic(err)
		}
		compare(t, strings.TrimSpace(string(stem)), Stem(string(word)), string(word))
	}
}
