package com.example.pariwarain.adapter

import androidx.fragment.app.Fragment
import androidx.fragment.app.FragmentActivity
import androidx.viewpager2.adapter.FragmentStateAdapter
import com.example.pariwarain.ui.properties.owned.OwnedPropertiesFragment
import com.example.pariwarain.ui.properties.rented.RentedPropertiesFragment

class PropertiesPagerAdapter(activity: FragmentActivity): FragmentStateAdapter(activity) {

    override fun createFragment(position: Int): Fragment {
        var fragment: Fragment? = null
        when (position) {
            0 -> fragment = RentedPropertiesFragment()
            1 -> fragment = OwnedPropertiesFragment()
        }
        return fragment as Fragment
    }

    override fun getItemCount(): Int {
        return 2
    }
}