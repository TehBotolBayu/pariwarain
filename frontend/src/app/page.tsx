import React from 'react'
import SearchBar from "@/components/SearchBar"
import Hero from '@/components/Hero'
import ProductList from '@/components/ProductList'
import Cardlist from '@/components/CardOffer'
import Guide from '@/components/Guide'
import Footer from '@/components/Footer'

function Home() {
  return (
    <div className=''>
      <SearchBar/>
      <Hero/>
      <h1 className="big-text">Nearby Property</h1>
      <ProductList />
      <h1 className="big-text">Most Popular</h1>
      <ProductList />
      <h1 className="big-text">What Could We Offer?</h1>
      <Cardlist/>
      <Guide/>
      <Footer/>
    </div>
  )
}

export default Home