import React from 'react'
import TipsText from './TipsText'
import Image from 'next/image'

function Guide() {
  return (
    <div className='guide-container'>
      <div>
        <h1 className='guide-title'>
          How To Do It?
        </h1>
        <Image
        src={'/billboard.png'}
        height={500}
        width={500}
        alt={'billboard'}
        className='h-full hidden md:block'
        />
      </div>
      <div className='text-center'>
        <TipsText text='List Your Property' h={''}/>
        <TipsText text='Create Canvas' h={''}/>
        <TipsText text='Make a Deal' h={''}/>
        <TipsText text='Monitor Your Ad' h={'h-0'}/>
      </div>
    </div>
  )
}

export default Guide