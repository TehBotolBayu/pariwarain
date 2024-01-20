import React from 'react'

function Card({
    title, 
    description
} : {
    title: string,
    description: string
}) {
  return (
    <div className='card'>
        <h1 className='text-lg font-bold'>{title}</h1>
        <p className='text-xs'>
            {description}
        </p>
    </div>
  )
}

export default Card