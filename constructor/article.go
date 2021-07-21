package constructor

type Article struct {
	Author       string
	Content      string
	Wordle       string //词云
	FirstPicture string //头图
	ReplyNum     int
	LikeNum      int
	Category     string
}

func (receiver *Article) Constuctor(author string, content string, firstPicture string) {
	receiver.Content = content
	receiver.Author = author
	receiver.Wordle = receiver.analyContent()
	if firstPicture == "" {
		receiver.FirstPicture = "https://img.php.cn/upload/article/000/000/039/5e040b6fa3b26130.jpg"
	} else {
		receiver.FirstPicture = firstPicture
	}
}
func (receiver Article) analyContent() (wordle string) {
	if receiver.Content != "xxx" {
		wordle = "XXXXXX"
		return
	}
	return
}
