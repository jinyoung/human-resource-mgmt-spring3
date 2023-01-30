package human.resource.mgmt.external;

import org.springframework.stereotype.Service;

@Service
public class CalendarServiceImpl implements CalendarService {

    public Calendar findByUserid(FindByUserIdQuery query) {
        Calendar calendar = new Calendar();
        return calendar;
    }

    public Calendar searchCalendar(SearchCalendarQuery query) {
        Calendar calendar = new Calendar();
        return calendar;
    }
}
