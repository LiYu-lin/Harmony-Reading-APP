import router from '@ohos.router'
import { Config } from './Index'
@Entry
@Component
struct Record {
 @StorageLink("user_config")
 config:Config={}

  build() {
Column(){
  Row() {
    Image($r('app.media.ic_public_arrow_left'))
      .width(30)
      .onClick(()=>{
        router.pushUrl({url:'pages/MainPage'})
      })
  }.height('5%')
  .backgroundColor(Color.White)

  .width("100%")
  Column({space:15}) {
ForEach(this.config.userInformation.records,(item:any,index:number)=>{
  RecordItem({Item:item})
})
      }.padding({top:15})


  }.backgroundColor(Color.Gray)

.height('100%')
    .width('100%')
  }
}
@Component
struct RecordItem{
Item:any
  build() {

      Column() {
        Row() {
          Row() {
            Text(this.Item.booKname)
            Text(this.Item.time.substring(0,10))



          }
          .width('100%')
          .justifyContent(FlexAlign.SpaceBetween)
          .padding({
            left: 15,
            right: 15
          })
          .borderRadius(8)
          .height(40)
          .backgroundColor(Color.White)
        }.padding({
          left: 10,
          right: 10
        })

      }
      .width('100%')

  }
}