package service

import (
	"fmt"
	"strings"

	"github.com/chillaso/wordealo/dao"
	"github.com/chillaso/wordealo/model"
)

const (
	BASE_LETTER_REGEXP = "(?=[a-z]*%s)"
	BASE_REGEXP        = "and palabra REGEXP '%s[a-z]+'\n"
	BASE_NOT_REGEXP    = "and palabra NOT REGEXP '%s[a-z]+'\n"
	BASE_LIKE          = "and palabra LIKE '%s'\n"
	BASE_NOT_LIKE      = "and palabra NOT LIKE '%s'\n"
	BASE_QUERY         = `SELECT DISTINCT palabra 
	FROM palabras
	WHERE tipo not in ('nombre/topónimo', 'preposición', 'pronombre', 'conjunción', 'otro/onomatopeya',
	'otro/locución', 'otro/expresión', 'otro/frase', 'otro/aumentativodiminutivo', 'otro/abreviatura',
	'otro/contracción', 'otro/interjección', 'artículo', 'nombre/propio')
	and tipo != ''
	and char_length(palabra) = %d
	%s LIMIT 15
	`
)

func GetPossibleWords(wordealo *model.Wordealo) ([]string, error) {

	correctQuery := getCorrectQuery(wordealo.Correct, wordealo.WordLength)
	incorrectQuery := getIncorrectQuery(wordealo.Incorrect, wordealo.WordLength)
	whereQuery := correctQuery + incorrectQuery

	query := fmt.Sprintf(BASE_QUERY, wordealo.WordLength, whereQuery)
	fmt.Println(query)

	result := dao.GetPossibleWords(query)
	fmt.Println(result)

	return result, nil
}

func getCorrectQuery(letters []model.Letter, wordLength int) string {
	likeCriteria := make([]string, wordLength)
	var regexpCriteria string
	var notLikeQuery string

	for _, letter := range letters {
		var isInRegexpCriteria = false
		for position, isRight := range letter.Position {
			if isRight {
				likeCriteria[position] = letter.Letter
			} else {
				if !isInRegexpCriteria {
					regexpCriteria += fmt.Sprintf(BASE_LETTER_REGEXP, letter.Letter)
					isInRegexpCriteria = true
				}

				notLikeCriteria := make([]string, wordLength)
				notLikeCriteria[position] = letter.Letter
				formatLikeCriteria(notLikeCriteria)
				notLikeQuery += fmt.Sprintf(BASE_NOT_LIKE, strings.Join(notLikeCriteria, ""))
			}
		}
	}

	formatLikeCriteria(likeCriteria)
	likeQuery := fmt.Sprintf(BASE_LIKE, strings.Join(likeCriteria, ""))
	regexpQuery := fmt.Sprintf(BASE_REGEXP, regexpCriteria)

	return likeQuery + regexpQuery + notLikeQuery
}

func getIncorrectQuery(letters []model.Letter, wordLength int) string {
	var notRegexpQuery string
	for _, letter := range letters {
		notRegexpCriteria := fmt.Sprintf(BASE_LETTER_REGEXP, letter.Letter)
		notRegexpQuery += fmt.Sprintf(BASE_NOT_REGEXP, notRegexpCriteria)
	}
	return notRegexpQuery
}

func formatLikeCriteria(criteria []string) {
	for i, s := range criteria {
		if s == "" {
			criteria[i] = "_"
		}
	}
}
