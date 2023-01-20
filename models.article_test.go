package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	allist := getAllArticles()

	if len(allist) != len(articleList) {
		t.Fail()
	}

	for i, v := range allist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}

}
