package io.github.antangelo.musync;

import io.github.antangelo.musync.status.StatusController;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.util.Assert;

@SpringBootTest
class MusyncApplicationTests
{
	@Test
	void contextLoads()
	{
	}

	@Test
	void statusTest()
	{
		Assert.hasLength(new StatusController().health());
	}
}
