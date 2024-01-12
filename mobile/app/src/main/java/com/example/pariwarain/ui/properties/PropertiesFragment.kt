package com.example.pariwarain.ui.properties

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import com.example.pariwarain.R
import com.example.pariwarain.adapter.PropertiesPagerAdapter
import com.example.pariwarain.databinding.FragmentPropertiesBinding
import com.google.android.material.tabs.TabLayoutMediator

class PropertiesFragment : Fragment() {

    private var _binding: FragmentPropertiesBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?,
    ): View {
        _binding = FragmentPropertiesBinding.inflate(inflater, container, false)
        val root: View = binding.root
        setupTabLayout()
        return root
    }

    private fun setupTabLayout() {
        val tabs = binding.tabs
        val viewPager = binding.viewPager
        val sectionsPagerAdapter = PropertiesPagerAdapter(requireActivity()) // Pass the activity to the adapter
        viewPager.adapter = sectionsPagerAdapter
        TabLayoutMediator(tabs, viewPager) { tab, position ->
            tab.text = resources.getString(TAB_TITLES[position])
        }.attach()
    }


    companion object {
        private val TAB_TITLES = intArrayOf(
            R.string.tab1,
            R.string.tab2
        )
    }
}