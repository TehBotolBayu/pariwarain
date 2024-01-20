import React from 'react'

function SearchBar() {
  return (
    <div className='bar-container'>
        <input type="text" placeholder='Search' className='searchbar' />
        <img src='/logo.png' className='absolute ml-5' width={50}/>
        <nav className='header-nav flex items-center'>
          <a href="">Papan Billboard</a>
          <a href="">Gedung</a>
          <a href="">Banner</a>
          <a href="">Hubungi Kami</a>
          <a href="">Daftar</a>
        </nav>
        <div className='md:hidden'>
          <img src='/hamburger-menu.svg'/>
        </div>
        <button className='button-see w-fit px-5 mx-5 hidden md:block'>
          Sign Up
        </button>
    </div>
  )
}

export default SearchBar