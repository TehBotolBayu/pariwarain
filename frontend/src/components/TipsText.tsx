import React from 'react'

function TipsText({
    text,
    h
}:{
    text: string,
    h: string
}) {
  return (
    <div className='flex ml-[30%] w-full'>
        <div className={`progress-line ${h}`}>
            <div className='progress-dot'></div>
        </div>
        <p className='text-white md:text-2xl -translate-y-3 ml-4'>
            {text}
        </p>
    </div>
  )
}

export default TipsText