package qtype

type QType int

const (
	SingleChoice         QType = iota // 单选题
	MultipleChoice                    // 多选题
	FillInTheBlank                    // 填空题
	TrueOrFalse                       // 判断题
	ShortAnswer                       // 简答题
	TermExplanation                   // 名词解释
	Essay                             // 论述题
	Calculation                       // 计算题
	QueOther                          // 其它
	JournalEntry                      // 分录题
	DocumentBased                     // 资料题
	Matching                          // 连线题
	Ordering                          // 排序题
	Cloze                             // 完型填空
	ReadingComprehension              // 阅读理解
	Oral                              // 口语题
	Listening                         // 听力题
	SharedOptions                     // 共用选项题
	Evaluation                        // 测评题
)

var QTypeStr = [...]string{
	"单选题",
	"多选题",
	"填空题",
	"判断题",
	"简答题",
	"名词解释",
	"论述题",
	"计算题",
	"其它",
	"分录题",
	"资料题",
	"连线题",
	"排序题",
	"完型填空",
	"阅读理解",
	"口语题",
	"听力题",
	"共用选项题",
	"测评题",
}

// 转QType转Str
func (q QType) String() string {
	return QTypeStr[q]
}

// Str转QType
func Index(qtypeStr string) QType {
	for k, v := range QTypeStr {
		if v == qtypeStr {
			return QType(k)
		}
	}
	return -1
}

func (q QType) YanXiString() string {
	return [...]string{
		"single",
		"multiple",
		"填空题",
		"judgement",
		"completion",
		"名词解释",
		"论述题",
		"计算题",
		"其它",
		"分录题",
		"资料题",
		"连线题",
		"排序题",
		"完型填空",
		"阅读理解",
		"口语题",
		"听力题",
		"共用选项题",
		"测评题",
	}[q]
}
