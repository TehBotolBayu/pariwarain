package com.example.pariwarain.ui.properties.rented

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.example.pariwarain.databinding.FragmentRentedPropertiesBinding

class RentedPropertiesFragment : Fragment() {

    private var _binding: FragmentRentedPropertiesBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?,
    ): View? {
        _binding = FragmentRentedPropertiesBinding.inflate(inflater, container, false)
        val root: View = binding.root
        return root
    }
}