import promptAction from '@ohos.promptAction'
import util from '@ohos.util'
import http from '@ohos.net.http'
import { Config,comment,api_addcomment } from './Index'

@Entry
@Component
struct ZhiHu {
  @StorageLink("user_config")
  config:Config={}
  addComment(item: comment) {
    this.config.books[this.config.index].comment.unshift(item)
  }

  changeLike(item: comment) {
    if (item.like_count) {
      item.like_count--
      item.likeFlag = false
    } else {
      item.like_count++
      item.likeFlag = true
    }
    let index = this.config.books[this.config.index].comment.findIndex(obj => obj.id === item.id)
    this.config.books[this.config.index].comment.splice(index, 1, item)
  }

  build() {
    Column() {
      Scroll() {
        Column() {
          // 顶部组件
          HmNavBar()
          // 评论组件
          // CommentItem({
          //   item: {
          //     id: 1,
          //     avatar: $r('app.media.ic_public_favor'),
          //     author: '周杰伦',
          //     content: '意大利拌面应该使用42号钢筋混凝土再加上量子力学缠绕最后通过不畏浮云遮望眼',
          //     time: '11-30',
          //     area: '海南',
          //     likeNum: 100
          //   }
          // })
          // // 分割线
          Divider()
            .strokeWidth(6)
            .color("#f4f5f6")
          ForEach(this.config.books[this.config.index].comment, (item:comment) => {
            CommentItem({
              item,
              changeLike: item => {
                this.changeLike(item)
              }
            })
          }, (item: comment) => JSON.stringify({ id: item.id, flag: item.likeFlag, num: item.like_count }))
        }
        .width('100%')
        .backgroundColor(Color.White)
      }.padding({
        bottom: 60
      })

      ReplyInput({
        addComment: item => {
          this.addComment(item)
        }
      })
    }.height('100%')
    .justifyContent(FlexAlign.SpaceBetween)
  }
}

@Component
struct HmNavBar {
  build() {
    Row() {
      Row() {
        Image($r('app.media.ic_public_arrow_left'))
          .width(20)
          .height(20)
      }
      .borderRadius(20)
      .backgroundColor('#f6f6f6')
      .justifyContent(FlexAlign.Center)
      .width(30)
      .aspectRatio(1)
      .margin({
        left: 15
      })

      Text("评论回复")
        .layoutWeight(1)
        .textAlign(TextAlign.Center)
        .padding({
          right: 35
        })
    }
    .width('100%')
    .height(50)
    .border({
      width: {
        bottom: 1
      },
      color: '#f6f6f6',
    })
  }
}

@Component
struct CommentItem {
  item:  Partial<comment> = {}
  changeLike: (item: comment) => void = () => {}

  build() {
    Row() {
      Image($r('app.media.ic_user_portrait'))
        .width(32)
        .height(32)
        .borderRadius(16)
      Column({ space: 10 }) {
        Text(this.item.user)
          .fontWeight(600)
        Text(this.item.content)
          .lineHeight(20)
          .fontSize(14)
          .fontColor("#565656")

        Row() {
          Text(`${this.item?.Create_time} . IP属地 ${"北京"}`)
            .fontColor("#c3c4c5")
            .fontSize(12)
          Row() {
            Image($r('app.media.ic_public_favor'))
              .width(14)
              .aspectRatio(1)
              .fillColor(this.item?.likeFlag ? 'red' : "#c3c4c5")

            Text(this.item?.like_count?.toString())
              .fontSize(12)
              .margin({
                left: 5
              })
          }.onClick(() => {
            this.changeLike(this.item as comment)
          })
        }
        .width('100%')
        .justifyContent(FlexAlign.SpaceBetween)

      }.alignItems(HorizontalAlign.Start)
      .layoutWeight(1)
      .padding({
        left: 15,
        right: 5
      })
    }.justifyContent(FlexAlign.Start)
    .alignItems(VerticalAlign.Top)
    .width('100%')
    .padding(15)
  }
}


@Component
struct ReplyInput {
  @StorageLink('user_config')
  config:Config={}
  @State commentStr: string = ''
  addComment: (item: comment) => void = () => {}
  async  Add(bit :string){
    try {


      const req=http.createHttp()
      const res= await req.request(api_addcomment+"bit="+bit)


      // let result:MessageChapter=JSON.parse( res.result as string) as MessageChapter
      // this.config.chapter_content=result.data
      // if(result.status=="success"){
      //   router.pushUrl({url:'pages/ReadPage'})
      // }
      // else{
      //   AlertDialog.show({message:"没有章节"})
      // }
      //console.log(JSON.stringify(this.config))
      // AlertDialog.show({message:JSON.stringify(this.config.chapter_content)})

    }catch (e){
      console.log("error")
      console.log(e)
    }
  }
  build() {
    Row() {
      TextInput({ placeholder: '回复', text: this.commentStr })
        .layoutWeight(1)
        .backgroundColor("#f4f5f6")
        .onChange(value => {
          this.commentStr = value
        })
      Text('发布')
        .fontColor("#6ecff0")
        .fontSize(14)
        .margin({
          left: 10
        })
        .onClick(() => {
          if (this.commentStr !== "") {
            let obj: comment = {
              id: this.config.books[this.config.index].id,
              content: this.commentStr,
              like_count: 0,
              Create_time:new(Date),
              child_Comments:[],
              //likeFlag: false,
              user: this.config.userInformation.name,
              //avatar: $r('app.media.ic_user_portrait'),
              //area: "北京"
            }
            this.addComment && this.addComment(obj)
            this.commentStr = ""
            this.Add(JSON.stringify(comment))

          }

        })
    }
    .height(50)
    .padding({
      left: 10,
      right: 10
    })
    .backgroundColor(Color.White)
    .border({
      width: { top: 1 },
      color: "#f4f5f6"
    })
  }
}


export class ReplyItem {
  id: number = 0
  avatar: string | Resource = ""
  author: string = ""
  content: string = ""
  time: string = ""
  area: string = ""
  likeNum: number = 0
  likeFlag?: boolean = false
}