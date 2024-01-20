import React from 'react'


function Footer() {
  return (
    <div className='footer-container justify-center '>
        <div className=''>
            <div>
                <h1 className='font-bold text-white text-lg mb-5'>Get Daily News About Our Service</h1>
            </div>
            <div className='flex bg-white rounded-md py-2 px-2 items-center mb-10'>
                <label htmlFor="sub" className='mr-2'><img src="/mail.svg" width={100}/></label>
                <input name='sub' type='text' placeholder='Enter Your Email'/>
                <button className='button-see text-base'>Subscribe</button>
            </div>
        </div>
        <div className='flex  md:justify-center md:w-1/2'>
            <div className='w-1/2 md:w-fit'>
                <img src='/logo.png' />
            </div>
            <div>
                <nav>
                    <ul className='text-white text-sm [&>li]:mb-2 md:flex [&>li]:md:mx-5'>
                        <li>Home</li>
                        <li>Billboard</li>
                        <li>Building</li>
                        <li>Houses</li>
                    </ul>
                </nav>
            </div>
        </div>

    </div>
  )
}

export default Footer