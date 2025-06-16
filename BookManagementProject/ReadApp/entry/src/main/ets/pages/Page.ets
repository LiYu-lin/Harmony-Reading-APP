import router from '@ohos.router'
import http from '@ohos.net.http'
import { api } from './Index'
class  Message {
  time: Date;
  type: string;
  status: string;
  data: null | any;
}
async function sleep(ms: number): Promise<void> {
  await new Promise((resolve) => setTimeout(resolve, ms));
}

const api_regist=api+'regist?'
@Entry
@Component
struct BuilderParamCard02 {
  @State message: string = 'Hello World'
  async getData() {
    try {
      const req = http.createHttp()

      const res = await  req.request(api_regist+"name="+this.name+
      "&password="+this.password+"&qq="+this.qq+"&phone="+this.phone+
      "&email="+this.email+"&introduce="+this.introduce)
      //console.log(res.result as string)
     let result:Message=JSON.parse(res.result as string)as Message
     //7 console.log(result)
     //  AlertDialog.show({
     //    message: JSON.stringify(result)
     //  })

      if (result.type == "regist" && result.status == "regist success") {
        AlertDialog.show({ message: "注册成功，返回登录" })
        await sleep(1000)
        router.pushUrl({ url: 'pages/Index' })
      } else if (result.status == "exist") {
        AlertDialog.show({ message: "用户名已存在，请更换" })
      } else if (result.status == "fail") {
        AlertDialog.show({ message: "信息不完整，请填写所有字段" })
      } else {
        AlertDialog.show({ message: result.status })
      }
      await sleep(1000)
      if(result.type=="regist"&&result.status=="regist success"){
        router.pushUrl({url:'pages/Index'})
      }
    } catch (e) {
      console.log('error')
    }
  }
  @State name: string=""
  @State password: string=""
  @State qq: string=""
  @State phone: string=""
  @State email: string=""
  @State introduce: string=""

  build() {
    Column({space:15}) {
      Row() {
        Image($r('app.media.ic_public_back'))
          .width(25)
          .aspectRatio(1)
          .onClick(()=>{
            router.pushUrl({
              url: '/pages/Index'
            })
          })
        Text('个人信息')
          .fontWeight(500)
          .fontSize(20)
        Text('保存')
          .fontColor(Color.Red)
          .onClick(()=>{
            this.getData()
            console.log('sf')
          })
      }.width('100%')
      .justifyContent(FlexAlign.SpaceBetween)
      .padding({left:10,right:10})
      Image($r('app.media.ic_user_portrait')).width(100)
      Text('用户头像')
      HmCardItem({ leftTitle: '姓名:', rightValue:$name})
      HmCardItem({leftTitle:'密码：',rightValue:$password})
      HmCardItem({ leftTitle: 'qq:', rightValue: $qq})
      HmCardItem({ leftTitle: '电话号码:', rightValue: $phone })
      HmCardItem({ leftTitle: '邮箱:', rightValue: $email })
      Introduce({ leftTitle: '自我介绍:', rightValue:$introduce })


    }
    .height('100%')
    .backgroundColor('#faf4f5f6')
  }
}
class Regist{
  name:string
  password:string
  qq:string
  phone:string
  email:string
  introduce:string
}
@Component
struct TopCard{
  build(){
    Row() {
      Image($r('app.media.ic_public_back'))
        .width(25)
        .aspectRatio(1)
        .onClick(()=>{
          router.pushUrl({ url: "pages/Index" })
        })
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
      .backgroundColor(Color.White)

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
  @Link
  rightValue: string
  leftTitle: string = ''

  build() {
    Row() {
      Text(this.leftTitle)
      TextInput({placeholder:this.rightValue}).borderRadius(0)
        .onChange((value)=>{
          this.rightValue=value
        })
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
  @Link
  rightValue: string
  build() {
    Column() {
      Text(this.leftTitle)
      TextArea({placeholder:this.rightValue}).fontColor("#ccc")
        .height(250)

    }.width('100%')
    .onAreaChange((value)=>{
      this.rightValue
    })
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

