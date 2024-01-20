import React from 'react'
import Card from './Card'

const offer = [
    {
        title:"Maximize Visibility",
        description:"From urban walls to subutban traces, our diverse range of advertising spaces ensures your brand is seen by the right audience"
    },
    {
        title:"Local Impact, Global Reach",
        description:"Reach your target audience where it matters most. Our platform bridges the gap between local spaces and global brands"
    },
    {
        title:"Maximize Visibility",
        description:"Track your campaigns performance effortlessly with our intuitive dashboard. Real-time updates, analytics and feedback empower you to make informed decision"
    },
    {
        title:"Trusted Partnership",
        description:"We foster trust between property owners and advertisers. Our transparent processes and verification systems ensure a seamless experience for all"
    }
]

function CardOffer() {
    

  return (
    <div className='offer-container'>
        <div className='cardoffer-container'>
        {
            offer.map((e, i) => {
                if(i%2!==0){
                    return (
                        <Card
                        title={e.title}
                        description={e.description}
                        key={i}
                        />
                    )
                }
            })
        }
        </div>
        <div className='cardoffer-container'>
        {
            offer.map((e, i) => {
                if(i%2===0){
                    return (
                        <Card
                        title={e.title}
                        description={e.description}
                        key={i}
                        />
                    )
                }
            })
        }
        </div>
    </div>
  )
}

export default CardOffer