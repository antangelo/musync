package io.github.antangelo.musync;

import io.github.antangelo.musync.status.StatusController;
import org.junit.Assert;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.mockito.junit.MockitoJUnitRunner;

@RunWith(MockitoJUnitRunner.class)
class MusyncApplicationTests
{
	@Test
	void testStatusPage()
	{
		Assert.assertEquals(new StatusController().health(), "OK");
	}
}
