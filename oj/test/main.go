package test

import (
	"baliance.com/gooxml/document"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//
//func main() {
//	//var filePath string
//	//flag.StringVar(&filePath, "fp", "", "pdf file path.")
//	//flag.Parse()
//	//if filePath == "" {
//	//	panic("file path must be provided")
//	//}
//	filePath := "C:\\oj\\test\\1.pdf"
//	content, err := ReadPdf(filePath) // Read local pdf file
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(content)
//
//	//将pdf的所有内容写入oem.html文件。(发现这个办法写文件很简单)
//	err = ioutil.WriteFile("./oem.html", []byte(content), 0777)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("保存成功")
//}
//
//func ReadPdf(path string) (string, error) {
//	f, err := os.Open(path)
//	defer f.Close()
//	if err != nil {
//		return "", err
//	}
//	//这个服务需要用docker生成：docker run -d -p 9998:9998 apache/tika:latest
//	client := tika.NewClient(nil, "http://127.0.0.1:9998")
//	return client.Parse(context.TODO(), f)
//}

func main() {
	wordIdentify("C:\\oj\\test\\1.docx", "C:\\oj\\test\\copy.docx")
	// 图片路径提供
	characterRecognition("C:\\Users\\zhoulinfeng\\Desktop\\html\\2.files")
}

// word 文字提取 源word 新word
func wordIdentify(wordPath, newWordPath string) {
	doc, _ := document.Open(wordPath)
	paragraphs := doc.Paragraphs()
	fp, _ := os.OpenFile(newWordPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm) // 读写方式打开
	// defer延迟调用
	defer fp.Close() //关闭文件，释放资源。
	// word 全部内容
	for _, paragraph := range paragraphs {
		//run为每个段落相同格式的文字组成的片段
		for _, run := range paragraph.Runs() {
			text := run.Text()
			fp.WriteString(text + "\t")
			fmt.Println(text)
		}
		fp.WriteString("\n")
	}
	//{  图片处理
	//log.Printf("转换完成")
	//for _, image := range doc.Images {
	//	//fmt.Println(image.Format())
	//	//fmt.Println(image.Path())
	//	fmt.Println(image.SetRelID)
	//}
	//log.Printf("图片解析完成")
	////邮箱
	//doc.MergeFields()
	//for i, image := range doc.Images {
	//	fmt.Println(i)
	//	fmt.Println(image.Path())
	//}}
	log.Printf("解析完成")
}

// 文字识别
func characterRecognition(path string) {
	// 图片解析追加文件
	fp, _ := os.OpenFile("C:\\oj\\test\\copy.docx", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm) // 读写方式打开
	defer fp.Close()
	dirs := Dirs(path)
	var stdout, stderr bytes.Buffer
	for _, figurePath := range dirs {
		cmd := exec.Command("tesseract", figurePath, "stdout", "-l", "chi_sim")
		cmd.Stdout = &stdout // 标准输出
		cmd.Stderr = &stderr // 标准错误
		err := cmd.Run()
		if err != nil {
			return
		}
		outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
		fp.WriteString(outStr)
		//log.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	}
}

func Dirs(path string) []string {
	// 所有目录
	var dirs []string
	// 所有图片路径
	var figurePaths []string
	// 获取图片目录地址 遍历
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("路径不存在：%v", err)
	}
	// "\\"
	separator := string(os.PathSeparator)
	// 遍历目录
	for _, info := range dir {
		// 判读是否是目录
		if info.IsDir() {
			dirs = append(dirs, path+separator+info.Name())
			Dirs(path + separator + info.Name())
		} else {
			// 图片
			figurePaths = append(figurePaths, path+separator+info.Name())
		}
	}
	return figurePaths
}
