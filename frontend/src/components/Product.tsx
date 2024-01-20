import React from 'react'
import Image from 'next/image'

function Product({
    img,
    price,
    address,
}:{
    img: string,
    price: number,
    address: string
}) {
  return (
    <div className='mx-3 md:mx-10'>
        <div className='product-container'>
            <Image
            src={img}
            height={100}
            width={100}
            alt={"property"}
            className='product-img'
            />
        </div>
        
        <div className='product-label '>
            <div className='font-bold text-xs'>Rp.{price}/hari</div>
            <div className='text-sm text-wrap'>{address}</div>
            <button className='button-see'>
                Lihat Detail
            </button>
        </div>
    </div>
  )
}

export default Product