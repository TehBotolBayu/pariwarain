package com.example.pariwarain.ui.properties.owned

import android.content.Intent
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.example.pariwarain.databinding.FragmentOwnedPropertiesBinding
import com.example.pariwarain.ui.properties.posting.PostingActivity

class OwnedPropertiesFragment : Fragment() {

    private var _binding: FragmentOwnedPropertiesBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?,
    ): View? {
        _binding = FragmentOwnedPropertiesBinding.inflate(inflater, container, false)
        val root: View = binding.root

        binding.btnPosting.setOnClickListener {
            startActivity(Intent(requireActivity(), PostingActivity::class.java))
        }

        return root
    }
}