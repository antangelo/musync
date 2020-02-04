package io.github.antangelo.musync.status;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
public class StatusController
{
	@RequestMapping("/health")
	@ResponseBody
	public String health()
	{
		return "OK";
	}
}
