package hw03frequencyanalysis

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{input: "some one", expected: []string{"some", "one"}},
		{input: "some,one", expected: []string{"some,one"}},
		{input: "some one ", expected: []string{"some", "one"}},
		{input: ",some one", expected: []string{",some", "one"}},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result := sliceStr(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string]int
	}{
		{input: []string{"some", "one"}, expected: map[string]int{"some": 1, "one": 1}},
		{input: []string{"some", "some", "one"}, expected: map[string]int{"some": 2, "one": 1}},
		{input: []string{"some", "one", "a", "", ""}, expected: map[string]int{"some": 1, "one": 1, "a": 1, "": 2}},
	}

	for _, tc := range tests {
		t.Run(tc.input[0], func(t *testing.T) {
			result := countWords(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestSortKeys(t *testing.T) {
	tests := []struct {
		input    map[string]int
		expected []string
	}{
		{input: map[string]int{"some": 1, "one": 2}, expected: []string{"one", "some"}},
		{input: map[string]int{"a": 2, "b": 3, "c": 1}, expected: []string{"b", "a", "c"}},
		{input: map[string]int{"a": 2, "b": 2, "c": 1}, expected: []string{"a", "b", "c"}},
		{input: map[string]int{"b": 2, "a": 2, "c": 1}, expected: []string{"a", "b", "c"}},
	}

	for _, tc := range tests {
		name := func(m map[string]int) string {
			var s strings.Builder
			for key := range m {
				s.WriteString(key)
			}
			return s.String()
		}(tc.input)
		t.Run(name, func(t *testing.T) {
			result := sortKeys(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestTrim(t *testing.T) {
	var data []string
	for r := 'a'; r <= 'z'; r++ {
		data = append(data, string(r))
	}
	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	result := trim(data, 10)
	require.Equal(t, expected, result)
}

// func TestTakeTop(t *tesing.T) {
//	tests := []struct {
//		input 	 map[string]int
//		expected []string
//	}{
//		{input: map[string]int{"some": 1, "one": 1}, expected: []string{}},
//	}
//}
