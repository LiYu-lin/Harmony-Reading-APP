import { Size } from '@ohos/hypium'
import router from '@ohos.router'
import { api_getchapter, chapter, Config,Book,chapter_content } from './Index'
import { MessageChapter } from './MainPage'
import http from '@ohos.net.http'
@Entry
@Component
struct Read{
  @StorageLink("user_config")
  config:Config={
}
  @State
  sun:boolean=true
  @State
  Size:number=1


  @State
  showSet:boolean=false
  @State
  showDirectory:boolean=false
  build(){
    Stack({alignContent:Alignment.Bottom})
    {
      ReadPage({showSet:$showSet,Size:$Size,sun:$sun})

if(this.showSet){
  Set({showSet:$showSet,Size:$Size,sun:$sun,showDirectory:$showDirectory})
  if(this.showDirectory)
  {
  Directory({showDirectory:$showDirectory,sun:$sun})
  }
}
    }
  }

}
@Component
struct Directory{
  @Link
  showDirectory :boolean
  @Link
  sun :boolean

  @StorageLink('user_config')
  config:Config={
  }
  async  GetChapter(id: number){
    try {


      const req=http.createHttp()
      const res= await req.request(api_getchapter+"id="+id.toString())

      let result:MessageChapter=JSON.parse( res.result as string) as MessageChapter
      this.config.chapter_content=result.data

      //console.log(JSON.stringify(this.config))
      // AlertDialog.show({message:JSON.stringify(this.config.chapter_content)})
      if(result.status=="success"){
        router.pushUrl({url:'pages/ReadPage'})
      }
      else {
        AlertDialog.show({ message: "没有章节" })
      }
    }catch (e){
      console.log("error")
      console.log(e)
    }
  }
  build() {

    Row(){
      Column() {
        Row() {
          Text(this.config.bookName)
            .fontSize(20)
            .padding({ top: 5 })
          //   List({space:10}){
          //     ForEach(this.config.books,(item:chapter,index:number)=>{
          //       ListItem(){
          //
          //      Dir({chap:item})
          //       }.padding({left:7,right:7})
          //       .onClick(()=>{
          //         this.config.bookName=item.name
          //         this.config.book_index=index
          //         this.config.chapter_id=item.id
          //         //this.GetChapter(item.chapters[0].id)
          //       })
          //     })
          // }

        }
          List({space:25}){
            ForEach(this.config.books[this.config.book_index].chapters,(item:chapter,index:number)=>{
              ListItem() {

                Dir({ chap: item })
              }.onClick(()=> {
                   // this.config.bookName=item.name
                   // this.config.book_index=index
                 this.config.chapter_id = item.id
                 this.GetChapter(item.id)
               })
            })
        }.height('100%')

      }
      .width('70%')
      .backgroundColor(this.sun ? '#fff6cd92' : '#999')
      Column(){

    }.height('100%')
      .width('30%')


      .onClick(()=>{
        this.showDirectory=!this.showDirectory
      })
    }.width("100%")
    .height('100%')
    .justifyContent(FlexAlign.Start)


  }
}

@Component
struct Dir{

 chap:chapter
build(){

Row(){
  Text(this.chap.name)
    .fontColor(Color.White)
    .fontSize(25)
    .fontWeight(600)
}.border({width:{bottom:1},color:Color.White})

}
}

@Component
struct ReadPage {
  @Link showSet:boolean
  @Link sun:boolean
  @Link Size:number
@StorageLink("user_config")
config:Config={}


  build() {
    Row() {
      Column() {

        Scroll(){
      Text(this.config.chapter_content.content||'sdfghjkjhgfdertyjuk激动飞机覅都是Jodi时间都手机电视降低送iv京东vi降低哦v道具i就DVD设计的降低哦v京东v觉得v哦道具i定居地vi就及vi跌幅巨大覅据覅v点击v点击v拒绝的sdfghjkjhgfdertyjuk激动飞机覅都是Jodi时间都手机电视降低送iv京东vi降低哦v道具i就DVD设计的降低哦v京东v觉得v哦道具i定居地vi就及vi跌幅巨大覅据覅v点击v点击v拒绝的sdfghjkjhgfdertyjuk激动飞机覅都是Jodi时间都手机电视降低送iv京东vi降低哦v道具i就DVD设计的降低哦v京东v觉得v哦道具i定居地vi就及vi跌幅巨大覅据覅v点击v点击v拒绝的sdfghjkjhgfdertyjuk激动飞机覅都是Jodi时间都手机电视降低送iv京东vi降低哦v道具i就DVD设计的降低哦v京东v觉得v哦道具i定居地vi就及vi跌幅巨大覅据覅v点击v点击v拒绝的\nsdfghjkjhgfdertyjuk激动飞机覅都是Jodi时间都手机电视降低送iv京东vi降低哦v道具i就DVD设计的降低哦v京东v觉得v哦道具i定居地vi就及vi跌幅巨大覅据覅v点击v点击v拒绝的sdfghjkjhgfdertyjuk激动飞机覅都是Jodi时间都手机电视降低送iv京东vi降低哦v道具i就DVD设计的降低哦v京东v觉得v哦道具i定居地vi就及vi跌幅巨大覅据覅v点击v点击v\n拒绝的')
        .fontSize(25/this.Size)
        .padding({left:5,right:5})
        .lineHeight(40/this.Size)
        .fontColor(this.sun?Color.Black:'#999')


      }
        .onClick(()=>{
          this.showSet=!this.showSet

        })
      }
      .width('100%')
      .height('100%')
      .backgroundColor(this.sun?'#fae3d2b4': '#ff292929')


    }
    .height('100%')
  }
}
@Component
struct  Set{
  @Link showSet:boolean
  @Link
  showDirectory :boolean
  @Link sun:boolean
  @Link Size:number
  @StorageLink("user_config")
  config:Config={}
  async  GetChapter(id: number){
    try {


      const req=http.createHttp()
      const res= await req.request(api_getchapter+"id="+id.toString())


      let result:MessageChapter=JSON.parse( res.result as string) as MessageChapter
      this.config.chapter_content=result.data
      if(result.status=="success"){
        router.pushUrl({url:'pages/ReadPage'})
      }
      else{
        AlertDialog.show({message:"没有章节"})
      }
      //console.log(JSON.stringify(this.config))
      // AlertDialog.show({message:JSON.stringify(this.config.chapter_content)})

    }catch (e){
      console.log("error")
      console.log(e)
    }
  }

  @StorageLink("user_config")
  Config :Config={}
  build(){
    Column(){
      Row(){
        Row() {
          Image($r('app.media.ic_public_arrow_left'))
            .width(35)
            .aspectRatio(1)

        }.padding({left:4,top:4})
       .onClick(()=>{
         router.pushUrl({
           url: 'pages/MainPage'
         })
       })
Text(this.Config.bookName)
  .fontSize(20)
  .padding({top:6,left:95})
      }.height("5%")
      .width('100%')
      .backgroundColor(Color.White)
      Row(){

      }.height("80%")
      .width('100%')

      .onClick(()=>{
        this.showSet=!this.showSet
      })
    Column(){
      Row() {
        if (this.sun) {
          Row() {
            Row() {
              Image($r('app.media.ic_public_donotdisturb'))
                .width(33)
            }.padding({ left: 5, right: 5, top: 5, bottom: 5 })
          }.borderRadius(300)
          .backgroundColor( '#fffccd8a')
          .onClick(() => {
            this.sun = !this.sun
          })

        } else {
          Row() {
            Row() {
              Image($r('app.media.ic_controlcenter_brightness_plus'))
                .width(33)
            }.padding({ left: 5, right: 5, top: 5, bottom: 5 })
          }.borderRadius(300)
          .backgroundColor('#ff474645')
          .onClick(() => {
            console.log("err1")
            this.sun = !this.sun
          })

        }
        Row(){
          Row(){
          Image($r('app.media.ic_public_comments'))
            .width(30)
          }
          .padding({ left: 5, right: 5, top: 5, bottom: 5 })
        }.borderRadius(300)
        .backgroundColor(this.sun ? '#fffccd8a' : '#ff474645')
        .onClick(() => {
         router.pushUrl({url:"pages/Comment"})

        })
      }.justifyContent(FlexAlign.SpaceBetween)
      .width('100%')
      .padding({left:5,right:5})
      Row(){
        Text('上一章')
          .onClick(()=>{
            this.config.chapter_id-=1

            this.GetChapter(this.config.chapter_id-1)

          })
        Text('A-')
        .onClick(()=>{
this.Size+=0.1
        })
        Column() {
          Image($r('app.media.ic_public_view_list'))
            .width(20)
        Text('目录')

        } .onClick(()=>{
          this.showDirectory=true
        })
        Text('A+')
          .onClick(()=>{
            this.Size-=0.1
          })
        Text('下一章').onClick(()=>{
          this.config.chapter_id+=1
          this.GetChapter(this.config.chapter_id)

        })


      }.justifyContent(FlexAlign.SpaceAround)
      .width('100%')
      .height(60)
      .backgroundColor(Color.White)
    }.height("15%")
    .width('100%')

    .justifyContent(FlexAlign.SpaceBetween)
  }.justifyContent(FlexAlign.SpaceBetween)
    .width('100%')
    .height('100%')
  }
}function build() {
throw new Error('Function not implemented.')
}
