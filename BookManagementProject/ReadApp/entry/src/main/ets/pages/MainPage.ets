import router from '@ohos.router'
import http from '@ohos.net.http'
import { api_getbook,Config,UserInformation,api_getuser,Book ,api_getchapter, chapter_content,chapter,MessageGetBooks,api_getrank,api_addrecord,MessageLogin} from './Index'
async function sleep(ms: number): Promise<void> {
  await new Promise((resolve) => setTimeout(resolve, ms));
}
const CATEGORY_RULES: Record<string, string[]> = {
  "玄幻": ["斗破", "修仙", "魔法", "大陆", "灵"],
  "言情": ["爱情", "恋", "甜", "心动"],
  "学习": ["Go", "语言", "开发", "教程", "编程"]
}
function filterBooksByCategory(books: Book[], category: string): Book[] {
  const keywords = CATEGORY_RULES[category]
  if (!keywords) return books

  return books.filter(book => {
    const text = (book.name + book.author + book.introduce).toLowerCase()
    return keywords.some(keyword => text.includes(keyword.toLowerCase()))
  })
}

@Entry
@Component
struct MainPage {
  @State message: string = 'Hello World'
  @State showuser:boolean=false
  build() {
    Row() {
      Column() {
       if(!this.showuser){
        Top()
       Main()}
        else{
userInformation()
        }



        Bottom({showUser:$showuser})
      }
      .width('100%')
      .height('100%')
      .justifyContent(FlexAlign.Start)
    }
    .height('100%')
  }
}
export class  MessageChapter{
  time: Date;
  type: string;
  status: string;
  data: chapter_content
}

@Component
struct Main{
  @StorageLink("user_config")
  config:Config={}
  @State
  Id :string=this.config.userInformation.id.toString()
  async AddRecord(bookname:string,id:string){
    try {
      const req=http.createHttp()
      const res= await req.request(api_addrecord+"id="+id+"&bookname="+bookname)

 let result:MessageLogin=JSON.parse( res.result as string) as MessageLogin
    if (result.status=="success") {

      this.config.userInformation=result.data

    }
    }catch (e){
      console.log(e)
    }
  }
  async GetChapter(id: number) {
    try {
      const req = http.createHttp()
      const res = await req.request(api_getchapter + "id=" + id.toString())

      console.log("📥 后端原始响应：", res.result)

      let result: MessageChapter = JSON.parse(res.result as string) as MessageChapter
      console.log("📤 解析后结果：", JSON.stringify(result))

      this.config.chapter_content = result.data

      if (result.status == "success") {
        router.pushUrl({ url: 'pages/ReadPage' })
      } else {
        AlertDialog.show({ message: "没有章节（后端状态：" + result.status + "）" })
      }
    } catch (e) {
      console.log("❌ 请求章节错误", e)
      AlertDialog.show({ message: "章节获取异常" })
    }
  }

  build(){
    List({space:10}){
      ForEach(this.config.books,(item:Book,index:number)=>{
        ListItem(){
          BookItem({item:item})
        }.padding({left:7,right:7})
        .onClick(()=>{
          this.AddRecord(item.name,this.Id)
          this.config.bookName=item.name
          this.config.book_index=index
          this.config.chapter_id=item.chapters[0].id
         this.GetChapter(item.chapters[0].id)

        })
      })

    }.backgroundColor('#fff8f1f1')
    .height("80%")


  }
}

export class bookItem {
  id:number=0
  name:string="侠名"
  author:string=""
  picture:string="https://ts1.cn.mm.bing.net/th/id/R-C.53347a63fc975020e7ce380c0f468c47?rik=0GYHbVaWZdLK4A&riu=http%3a%2f%2fimg95.699pic.com%2fphoto%2f50036%2f9962.jpg_wh860.jpg&ehk=B2wiOlzaLuSEseRhptzicOyaoJzoB8S%2bECDetjPPwD4%3d&risl=&pid=ImgRaw&r=0"
  like:number=0
  introduce:string=" 帮助读者使用相关Go语言资源快速开发自己的谷歌程序；充分利用Go的轻量级运行和并发特性一步步展示如何开发定制应用程序；展示后端开发中的编程过程和功能、详细示例、代码示例和节省时间的技巧。"
}
@Component
struct BookItem{
  item:Partial<Book>
  build(){
    Row({space:15}){
      Row() {
        Image(this.item.picture||"https://ts1.cn.mm.bing.net/th/id/R-C.53347a63fc975020e7ce380c0f468c47?rik=0GYHbVaWZdLK4A&riu=http%3a%2f%2fimg95.699pic.com%2fphoto%2f50036%2f9962.jpg_wh860.jpg&ehk=B2wiOlzaLuSEseRhptzicOyaoJzoB8S%2bECDetjPPwD4%3d&risl=&pid=ImgRaw&r=0")
          .width(80)
          .aspectRatio(0.65)
          .backgroundColor(Color.Blue)
          .borderRadius(13)
      }
      .padding({left:10,top:8,bottom:8})

Column({space:10}) {
  Row() {
    Text(this.item.name||"斗破苍穹")
      .fontSize(20)
      .fontWeight(600)
    Row() {
      Text('6.6分')
        .padding({left:5,right:5})
        .fontSize(13)
        .fontColor('#ffbf9103')
    }.borderRadius(13)
    .backgroundColor('#fffceccd')
  }.width('68%')

  .justifyContent(FlexAlign.SpaceBetween)
  Text(this.item.author||"侠名")
    .fontSize(15)
    .fontWeight(300)
  Text(this.item.Hit+"热度")
    .fontColor(Color.Red)
    .fontSize(13)
Text(this.item.introduce||" 帮助读者使用相关Go语言资源快速开发自己的谷歌程序；充分利用Go的轻量级运行和并发特性一步步展示如何开发定制应用程序；展示后端开发中的编程过程和功能、详细示例、代码示例和节省时间的技巧。")
  .textOverflow({
    overflow:TextOverflow.Ellipsis
  })
  .maxLines(2)
  .width('70%')
  .fontColor('#333')

}.alignItems(HorizontalAlign.Start)

    }
    .backgroundColor(Color.White)
    .borderRadius(15)

  }
}
@Component
struct Top {

  @StorageLink("user_config")
  config: Config = {}

  @State
  SearchBookName: string = ""

  @State
  list: string[] = ["精选", "玄幻", "学习", "言情", "排行"]

  @State
  selectIndex: number = this.config.index

  @Builder
  NatItem(isClick: boolean[], str: string, index: number) {
    Column() {
      Text(str)
        .fontSize(isClick ? 29 : 22)
        .fontWeight(400)
        .fontColor(Color.White)
    }
  }

  async GetBook(api: string) {
    try {
      const req = http.createHttp()
      const res = await req.request(api)
      console.log(res.result as string)
      let result: MessageGetBooks = JSON.parse(res.result as string) as MessageGetBooks
      this.config.books = result.data
      if (api == api_getrank) {
        await sleep(1000)
      }
      router.pushUrl({ url: 'pages/MainPage' })
    } catch (e) {
      console.log(e)
    }
  }

  // 新增：根据书名关键词搜索书籍
  async GetBookByName(name: string) {
    if (!name) {
      console.log("搜索关键词不能为空")
      return
    }
    const api_searchbook = "http://26.117.97.63:8080/book/searchbook?"
    const apiSearch = api_searchbook + "bookname=" + encodeURIComponent(name)
    console.log("调用搜索接口URL：" + apiSearch)
    await this.GetBook(apiSearch)
  }

  build() {
    Column({ space: 15 }) {
      Row({ space: 15 }) {
        Row() {
          Image($r('app.media.ic_public_input_search'))
            .width(25)
            .aspectRatio(1)
        }.padding({ left: 10 })

        TextInput({ placeholder: '搜索书名，作者，一搜既有' })
          .fontColor(Color.Gray)
          .layoutWeight(2)
          .onChange((value) => {
            this.SearchBookName = value
          })

        Button("搜索")
          .layoutWeight(1)
          .onClick(() => {
            if (this.SearchBookName && this.SearchBookName.trim() !== "") {
              console.log("搜索关键词: " + this.SearchBookName)
              this.GetBookByName(this.SearchBookName)
            } else {
              console.log("请输入搜索关键词")
            }
          })
      }
      .backgroundColor(Color.White)
      .width('90%')
      .height('39%')
      .borderRadius(15)

      Row() {
        ForEach(this.list, (item: string, index: number) => {
          Text(item)
            .fontSize(this.selectIndex === index ? 29 : 22)
            .fontWeight(400)
            .fontColor(Color.White)
            .onClick(async () => {
              this.selectIndex = index
              this.config.index = index

              if (item === '排行') {
                this.GetBook(api_getrank)
              } else if (item === '精选') {
                this.GetBook(api_getbook)
              } else {
                // 获取所有书籍后手动筛选
                const req = http.createHttp()
                const res = await req.request(api_getbook)
                const result: MessageGetBooks = JSON.parse(res.result as string) as MessageGetBooks

                const allBooks = result.data
                const filteredBooks = filterBooksByCategory(allBooks, item)
                this.config.books = filteredBooks
                router.pushUrl({ url: 'pages/MainPage' })
              }
            })


        })
      }
      .justifyContent(FlexAlign.SpaceAround)
      .width('100%')
    }
    .backgroundColor('#fa65a3f6')
    .width('100%')
    .height('13%')
    .padding({ top: 15 })
    .onClick(() => { })
  }
}

@Component
struct Bottom{
  @Link showUser:boolean
@Builder
bottomItem(image:string,name:string){
  Column({space:3}) {
    Image($r(image))
      .height(30)
      .aspectRatio(1)
  Text(name).fontSize(13)
    .fontColor('#ffb0acac')
    .fontWeight(800)
  }


}
  build(){
Row(){

  Column({space:3}) {
    Image($r('app.media.ic_gallery_frame_overlay_rectify'))
      .height(30)
      .aspectRatio(1)
    Text('书城').fontSize(13)
      .fontColor('#ffb0acac')
      .fontWeight(800)
  }.onClick(()=>{
    this.showUser=false
  })
  Column({space:3}) {
    Image($r('app.media.ic_device_foldablescreen'))
      .height(30)
      .aspectRatio(1)
    Text('阅读').fontSize(13)
      .fontColor('#ffb0acac')
      .fontWeight(800)

  } .onClick(()=>{
    router.pushUrl({
      url:'pages/ReadPage'})


  })
  Column({space:3}) {
    Image($r('app.media.ic_contacts_business_cards'))
      .height(30)
      .aspectRatio(1)
    Text('我的').fontSize(13)
      .fontColor('#ffb0acac')
      .fontWeight(800)
  }.onClick(()=>{
    this.showUser=true
  })

}.width('100%')
    .height('7%')
    .backgroundColor(Color.White)
.justifyContent(FlexAlign.SpaceAround)
    .border({width:{top:0.5,bottom:0.5},color:'#999'})
  }
}
@Component
struct userInformation{
  @StorageLink("user_config")
  config:Config={}
  Id=this.config.userInformation.id
  async Getinfor(){
    try {
      const req=http.createHttp()
      const res= await req.request(api_getuser+"id="+this.Id.toString())
      //console.log(res.result as string)
      let result:MessageLogin=JSON.parse( res.result as string) as MessageLogin

        this.config.userInformation=result.data

      console.log(JSON.stringify(this.config.userInformation))


    }catch (e){
      console.log(e)
    }
  }

  name:string=this.config.userInformation.name
  qq:string=this.config.userInformation.qq
  phone:string=this.config.userInformation.phone
  email:string=this.config.userInformation.email
  introduce:string=this.config.userInformation.introduce

  build(){
    Column({space:10}) {
      Image($r('app.media.ic_user_portrait')).width(100)
      Text('用户头像')
Column(){
  Image($r('app.media.ic_public_history')).width(20)
  Text('历史记录')

}.onClick(()=>{
  this.Getinfor()
  router.pushUrl({
    url: 'pages/Record'
  })}

)
      HmCardItem({ leftTitle: '姓名:', rightValue:this.name })
      HmCardItem({ leftTitle: 'qq:', rightValue: this.qq })
      HmCardItem({ leftTitle: '电话号码:', rightValue: this.phone })
      HmCardItem({ leftTitle: '邮箱:', rightValue: this.email})
      Introduce({ leftTitle: '自我介绍:', rightValue: this.introduce })
    }
    .height('93%')
      .backgroundColor('#faf4f5f6')
  }
}
@Component
struct TopCard{
  build(){
    Row() {
      Image($r('app.media.ic_public_back'))
        .width(25)
        .aspectRatio(1)
      Text('个人信息')
        .fontWeight(500)
        .fontSize(20)
      Text('保存')
        .fontColor(Color.Red)
        .onClick(()=>{

        })
    }.width('100%')
    .justifyContent(FlexAlign.SpaceBetween)
    .padding({left:10,right:10})
  }

}
@Component
struct HmCard {
  @BuilderParam
  CardFn: () => void
  build() {
    Column() {
      Column() {
        if(this.CardFn) {
          this.CardFn()
        }
      }.borderRadius(8)
      .backgroundColor('#F7F9FF')

    }.padding({
      left: 15,
      right: 15
    })
    .margin({
      top: 10
    })


  }
}
@Component
struct HmCardItem {
  leftTitle: string = ''
  rightValue: string = ''
  build() {
    Row() {
      Text(this.leftTitle)
      TextInput({placeholder:this.rightValue}).borderRadius(0)
        .width('65%')
    }.width('100%')
    .justifyContent(FlexAlign.SpaceBetween)
    .padding({
      left: 10,
      right:10
    })
    .height(50)
    .border({
      width: {
        bottom: 1
      },
      color: '#f4f5f6'
    })
    .backgroundColor(Color.White)
  }



}
@Component
struct Introduce {
  leftTitle: string = ''
  rightValue: string = ''
  build() {
    Column() {
      Text(this.leftTitle)
      TextArea({placeholder:this.rightValue}).fontColor("#ccc")
        .height(250)

    }.width('100%')
    .justifyContent(FlexAlign.SpaceBetween)
    .padding({
      left: 10,
      right: 10
    })

    .border({
      width: {
        bottom: 1
      },
      color: '#fffcfbfb'
    })
  }



}

