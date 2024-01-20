import React from 'react'
import Product from './Product'

const data = [
    {
        img: '/adboard.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/Rectangle 42.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/Rectangle 44.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/Rectangle 45.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/adboard.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/Rectangle 42.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/Rectangle 44.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/Rectangle 45.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/adboard.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/Rectangle 42.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/Rectangle 44.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
    {
        img: '/Rectangle 45.png',
        price: 10000,
        address: 'Kemayoran, Jakarta'
    },
]

function ProductList() {
  return (
    <div className='max-w-screen-xl md:mx-10'>
        <div className='list-container'>
        {
            data.map((e, i) => 
            <Product
            img={e.img}
            price={e.price}
            address={e.address}
            key={i}
            />
            )
        }
        </div>
    </div>
  )
}

export default ProductList