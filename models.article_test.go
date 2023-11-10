package main
import "testing"
func TestGetAllArticles(t *testing.T){
	alist := getAllarticles()
	if len(alist)!=len(articleList) { //checking for length
		t.Fail()
	}
	for i,v := range alist {//verifying the data obtained with the original data
		if v.Content!=articleList[i].Content || v.ID!=articleList[i].ID || v.Title!=articleList[i].Title{
			t.Fail();
			break;
		}
	}
}
