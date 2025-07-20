package site

// 根据枫枫的页面内容，设计了7个大类  // yaml:"这里要用小驼峰命名法"
type SiteInfo struct { // 站点信息
	Title string `yaml:"title" json:"title"`                   // 网站的title---》也就是浏览器顶端显示的那个标题
	Logo  string `yaml:"logo" json:"logo"`                     // 站点logo
	Beian string `yaml:"beian" json:"beian"`                   // 备案号(编号)
	Mode  int8   `yaml:"mode" json:"mode" binding:"oneof=1 2"` // 站点模式，1为社区模式 2为博客模式  // 我尼玛，这 binding:"required" 一定一定不要有空格，我真你妈服了，有空格会报错
}

type Project struct { // 项目设置
	Title   string `yaml:"title" json:"title"`     // 网站的标题
	Icon    string `yaml:"icon" json:"icon"`       // 项目图标
	WebPath string `yaml:"webPath" json:"webPath"` // 前端地址
}
type Seo struct { // SEO设置
	Keywords    string `yaml:"keywords" json:"keywords"`       // 关键词
	Description string `yaml:"description" json:"description"` // 网站描述

}
type About struct { // 关于我们
	SiteDate string `yaml:"siteDate" json:"siteDate"` // 网站创建时间 年月日
	QQ       string `yaml:"qq" json:"qq"`             // QQ的链接及二维码
	Version  string `yaml:"-" json:"version"`         // 网站版本号 不要读配置文件
	Wechat   string `yaml:"wechat" json:"wechat"`     // 微信的链接及二维码
	Gitee    string `yaml:"gitee" json:"gitee"`       // 码云的链接
	Bilibili string `yaml:"bilibili" json:"bilibili"` // 哔哩哔哩的链接
	Github   string `yaml:"github" json:"github"`     // Github的链接

}

type Login struct { // 登录设置
	QQLogin          bool `yaml:"qqLogin" json:"qqLogin"`                   // 使用qq登录
	UsernamePwdLogin bool `yaml:"usernamePwdLogin" json:"usernamePwdLogin"` // 使用用户名密码登录
	EmailLogin       bool `yaml:"emailLogin" json:"emailLogin"`             // 使用邮箱登录
	Captcha          bool `yaml:"captcha" json:"captcha"`                   // 启用登录验证码
}

type ComponentInfo struct { // 组件信息
	Title  string `yaml:"title" json:"title"`   // 组件的标题 ---》也就是 首页右侧组件展示 中的那些组件
	Enable bool   `yaml:"enable" json:"enable"` // 是否启用该组件
}

/*
因为排序操作涉及到对多个元素进行比较和重新排列。
列表（或数组）是一种非常适合这种操作的数据结构，因为它可以存储多个元素，并且可以方便地访问和修改这些元素的位置。
列表允许动态地添加、删除和修改组件。
因此，我们可以将组件信息配置为一个列表，并在首页右侧组件展示中，根据配置的顺序，动态地展示这些组件。
*/
type IndexRight struct { // 首页右侧组件展示（这里面 需要既能启用又能排序---》所以配置必须是一个列表，所以还需要继续套）
	List []ComponentInfo `json:"list" yaml:"list"` // 组件列表

}
type Article struct { // 文章设置
	NoExamine   bool `json:"noExamine" yaml:"noExamine" `    // 不需要审核文章
	CommentLine int  `json:"commentLine" yaml:"commentLine"` // 评论的层级
}

// ......................................
