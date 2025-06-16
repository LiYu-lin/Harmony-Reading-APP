import data from '@ohos.telephony.data'
import http from '@ohos.net.http'
import router from '@ohos.router'
import intl from '@ohos.intl'
import { expect } from '@ohos/hypium'
import { bookItem } from './MainPage'
export const ip="26.117.97.63"
export const port="8080"
export const api=ip+':'+port+'/user/'
export const api_getuser=api+"getinformation?"
export const api_getbook=ip+':'+port+'/book/getbook?'
export const api_getchapter=ip+':'+port+'/book/getchapter?'
export const api_getrank=ip+':'+port+'/book/getrank?'
const api_login=api+'login?'
export const api_addrecord=ip+':'+port+'/user/addrecord?'
export const api_addcomment=ip+':'+port+'/book/addcomment?'
PersistentStorage.PersistProp("user_config",'')

export class chapter_content{
  id:number
  book_id:number
  chapter_name:string
  chapter_old :number
  create_time:Date
  update_time:Date
  content:string
}
export  class Config{
  size?:number

  sun?:boolean
  showDirectory?:boolean
  bookName?:string
  directory?:Directory[]
  userInformation?:UserInformation
  books?:Book[]
  chapter_content?:chapter_content
  book_index?:number
  chapter_id?:number
  index?:number
}

let configInit:Config={
size:1,
  sun:true,
  showDirectory:false,
  bookName:"",
  directory:[],
  userInformation:null,
  books:null,
  chapter_content:null,
  index:0

}
export class UserInformation{
  id :number
  name:string
  password:string
  qq:string
  phone:string
  email:string
  introduce:string
  look_record:Date
  regist_time:Date
  lastest_login_time:Date
  age:string
  records:any
}
export class  MessageLogin {
  time: Date;
  type: string;
  status: string;
  data: UserInformation;
}
export class MessageGetBooks{
  time: Date;
  type: string;
  status: string;
  data:Book[];
}
export class Book{
  id:number
  create_time:string
  update_time:string
  author:string
  chapters_count:number
  like_count:number
  comment:comment[]
  name:string
  look_count:number
  Hit:number
  chapters:chapter[]
  picture?:string
  introduce?:string


}
export class chapter{
id :number
  name:string
}
export class comment{
  id:number
  content:string
  Create_time:Date
  user:string
  child_Comments:comment[]
  like_count :number
  likeFlag?:boolean
}

@Entry
@Component
struct Index {
  @State message: string = 'Hello World'

  build() {

    Column() {
      login()
        .padding({left:20,right:20})

    }
    .justifyContent(FlexAlign.Center)
    .backgroundColor(Color.White)
    .width('100%')
    .height('100%')

  }
}

class Directory{
  name:string
  id:number
}
async function sleep(ms: number): Promise<void> {
  await new Promise((resolve) => setTimeout(resolve, ms));
}
@Component
struct login{
  @StorageLink("user_config")
  config:Config={}
  @State name:string=''
  @State password:string=''
  @State isLogin :boolean=false
  async GetBook(){
    try {


      const req=http.createHttp()
      const res= await req.request(api_getbook)
      console.log(res.result as string)
      let result:MessageGetBooks=JSON.parse( res.result as string) as MessageGetBooks
     this.config.books=result.data
     // AlertDialog.show({message:JSON.stringify(this.config.books)})

    }catch (e){
      console.log(e)
    }
  }
async Login(){
  try {
    const req=http.createHttp()
    const res= await req.request(api_login+"name="+this.name+"&password="+this.password)
    //console.log(res.result as string)
    let result:MessageLogin=JSON.parse( res.result as string) as MessageLogin
    if (result.status=="success"&&result.type=="login") {
      InitConfig()
      AlertDialog.show({ message: "登录成功" })
      this.GetBook()
     // console.log(JSON.stringify(result))

      this.config.userInformation=result.data
      //console.log(JSON.stringify(this.config))

      await sleep(1000)
      router.pushUrl({ url: "pages/MainPage" })
    }else{
      AlertDialog.show({ message: "登录失败:请检查账号密码"+result.status })
    }

  }catch (e){
console.log(e)
  }
  }

  build(){
    Column({space:15}) {
      Row() {
        Row() {
          Image($r('app.media.ic_user_portrait'))
            .width(30)
            .aspectRatio(1)


        }
        .width(50)
        .height(50)
        .backgroundColor(Color.White)
        .justifyContent(FlexAlign.Center)
        .borderWidth(2)
        .borderColor('#ffa4a4a3')
        Row() {
          TextInput({ placeholder: 'Username' })
            .backgroundColor(Color.White)
            .onChange(value => {
              this.name = value
              if (this.name !== "" && this.password !== "") {
                this.isLogin = true
              }
              else {
                this.isLogin = false
              }
            })
            .borderRadius(1)
            .height(47)
        }
        .layoutWeight(1)
        .border({width:{bottom:2,right:2,top:2}})
        .borderColor('#ffa4a4a3')

      }.width('100%')
      Row() {
        Row() {
          Image($r('app.media.ic_public_lock'))
            .width(30)
            .aspectRatio(1)


        }
        .width(50)
        .height(50)
        .backgroundColor(Color.White)
        .justifyContent(FlexAlign.Center)
        .borderWidth(2)
        .borderColor('#ffa4a4a3')
        Row() {
          TextInput({ placeholder: 'Password' })
            .type(InputType.Password)
            .onChange(value => {
              this.password=value
              if (this.name !== "" && this.password !== ""){
                this.isLogin=true
              }
              else{
                this.isLogin=false
              }
            })
            .backgroundColor(Color.White)

            .borderRadius(1)
            .height(47)
        }
        .layoutWeight(1)
        .border({width:{bottom:2,right:2,top:2}})
        .borderColor('#ffa4a4a3')

      }.width('100%')

      Row() {
        Text("register now")
          .fontColor('#ff4982f1')
        .onClick(()=>{
          router.pushUrl({url:'pages/Page'})
        })
        Row() {
          Button("LOGIN")
            .onClick((event: ClickEvent) => {
this.Login()
            })
            .fontWeight(800)
            .width('30%')
            .aspectRatio(2)
            .enabled(this.isLogin)
            .type(ButtonType.Normal)
        }
      }.width('100%')
      .justifyContent(FlexAlign.SpaceBetween)

    }

  }

}

 export function InitConfig(){
  AppStorage.SetOrCreate<Config>("user_config",configInit)
 }
export function GetConfig(){
 return AppStorage.Get<Config>("user_config")

}
export function  Altercongig(newConfig:Config){


AppStorage.Set<Config>("user_config",newConfig)
}